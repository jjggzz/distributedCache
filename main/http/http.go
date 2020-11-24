package http

import (
	"Fool/main/cache"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
	log.Print("开始监听服务...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

type cacheHandler struct {
	*Server
}

func (h *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 获取key
	key := strings.Split(r.RequestURI, "/")[2]
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	m := r.Method
	// get
	if m == http.MethodGet {
		bytes, err := h.Get(key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(bytes) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_, _ = w.Write(bytes)
		return
	}

	// set
	if m == http.MethodPut {
		bytes, _ := ioutil.ReadAll(r.Body)
		if len(bytes) != 0 {
			err := h.Set(key, bytes)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
		return
	}

	// del
	if m == http.MethodDelete {
		err := h.Del(key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)

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
