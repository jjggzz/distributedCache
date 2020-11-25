package tcp

import "distributeCache/tcpimpl/cache"

func New(c cache.Cache) *Server {
	return &Server{c}
}
