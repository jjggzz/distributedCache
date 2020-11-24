package http

import (
	"io/ioutil"
	"net/http"
	"strings"
)

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
