package http

import (
	"Fool/main/cache"
	"encoding/json"
	"log"
	"net/http"
)

func New(c cache.Cache) *Server {
	return &Server{c}
}

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status", s.statusHandler())
}

type cacheHandler struct {
	*Server
}

func (h *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO 缓存操作方法未实现
}

func (s *Server) cacheHandler() http.Handler {
	return &cacheHandler{s}
}

type statusHandler struct {
	*Server
}

func (h *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	bytes, err := json.Marshal(h.GetStat())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(bytes)
}

func (s *Server) statusHandler() http.Handler {
	return &statusHandler{s}
}
