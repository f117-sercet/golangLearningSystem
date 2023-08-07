package lru

import "container/list"

type Cache struct {
	maxBytes  int64
	nobytes   int64
	li        *list.List
	cache     map[string]*list.Element
	OnEvicted func(key string, value Value)
}

// 计算使用了多少bit
type Value interface {
	Len() int
}
type entry struct {
	key   string
	Value string
}
