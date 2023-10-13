package rpcFramework

import (
	"encoding/json"
	"golangLearningSystem/goproject/go_code/goFramework/rpcFramework/codec"
	"io"
	"log"
	"net"
	"sync"
)

const MagicNumber = 0x3bef5c

type Option struct {
	MagicNumber int
	CodecType   codec.Type
}

var DefaultOption = &Option{

	MagicNumber: MagicNumber,
	CodecType:   codec.GobType,
}

type Server struct{}

func NewServer() *Server {

	return &Server{}
}

func (server *Server) ServerConn(conn io.ReadWriteCloser) {

	defer func() { _ = conn.Close() }()
	var opt Option
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {

		log.Println("rpc server:options err:", err)
		return

	}

	if opt.MagicNumber != MagicNumber {
		log.Printf("rpc server: invalid magic number %x", opt.MagicNumber)
		return
	}
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		log.Printf("rpc server: invalid codec type %s", opt.CodecType)
		return
	}
	server.serveCodec(f(conn))
}

var invalidRequest = struct {
}{}

func (server *Server) serverCodec(cc codec.Codec) {

	seding := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	for {
		req, err := server.readRequest(cc)
		if err != nil {
			if req == nil {
				break
			}
			req.h.Error = err.Error()
			server.sendResponse(cc, req.h, invalidRequest, sending)
		}
	}

}

func Accept(lis net.Listener) { DefaultServer.Accept(lis) }
