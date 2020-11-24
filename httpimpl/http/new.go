package http

import (
	"distributeCache/httpimpl/cache"
)

func New(c cache.Cache) *Server {
	return &Server{c}
}
