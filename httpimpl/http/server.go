package http

import (
	"distributeCache/httpimpl/cache"
	"fmt"
	"log"
	"net/http"
)

// 缓存服务类
type Server struct {
	cache.Cache
}

// 监听http请求
func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status", s.statusHandler())
	log.Print("开始监听服务...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Server) cacheHandler() http.Handler {
	return &cacheHandler{s}
}

func (s *Server) statusHandler() http.Handler {
	return &statusHandler{s}
}
