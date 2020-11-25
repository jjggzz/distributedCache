package tcp

import (
	"distributeCache/tcpimpl/cache"
	"net"
)

// 缓存服务类
type Server struct {
	cache.Cache
}

// 监听tcp请求
func (s *Server) Listen() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		// TODO 开一个协程去处理
	}
}
