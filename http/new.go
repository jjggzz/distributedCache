package http

import (
	"distributeCache/cache"
)

func New(c cache.Cache) *Server {
	return &Server{c}
}
