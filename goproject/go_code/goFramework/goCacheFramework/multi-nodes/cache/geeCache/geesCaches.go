package geeCache

import (
	"fmt"
	"golangLearningSystem/goproject/go_code/goFramework/goCacheFramework/multi-nodes/cache"

	//"golangLearningSystem/goproject/go_code/goFramework/goCacheFramework/multi-nodes/cache/lru"
	"log"
	"sync"
)

type Group struct {
	name      string
	getter    Getter
	mainCache cache.Cache
	peers     cache.PeerPicker
}
type Getter interface {
	Get(key string) ([]byte, error)
}
type GetterFunc func(key string) ([]byte, error)

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache.Cache{CacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}

func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}
func (g *Group) Get(key string) (cache.ByteView, error) {
	if key == "" {
		return cache.ByteView{}, fmt.Errorf("key is required")
	}

	if v, ok := g.mainCache.Get(key); ok {
		log.Println("[GeeCache] hit")
		return v, nil
	}

	return g.load(key)
}

func (g *Group) RegisterPeers(peers cache.PeerPicker) {
	if g.peers != nil {
		panic("RegisterPeerPicker called more than once")
	}
	g.peers = peers
}
func (g *Group) load(key string) (value cache.ByteView, err error) {
	if g.peers != nil {
		if peer, ok := g.peers.PickPeer(key); ok {
			if value, err = g.getFromPeer(peer, key); err == nil {
				return value, nil
			}
			log.Println("[GeeCache] Failed to get from peer", err)
		}
	}

	return g.getLocally(key)
}

func (g *Group) populateCache(key string, value cache.ByteView) {
	g.mainCache.Add(key, value)
}

func (g *Group) getLocally(key string) (cache.ByteView, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return cache.ByteView{}, err

	}
	value := cache.ByteView{B: cache.CloneBytes(bytes)}
	g.populateCache(key, value)
	return value, nil
}

func (g *Group) getFromPeer(peer cache.PeerGetter, key string) (cache.ByteView, error) {
	bytes, err := peer.Get(g.name, key)
	if err != nil {
		return cache.ByteView{}, err
	}
	return cache.ByteView{B: bytes}, nil
}
