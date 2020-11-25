package main

import (
	"distributeCache/httpimpl/cache"
	"distributeCache/httpimpl/http"
)

func main() {
	http.New(cache.New("inMemory")).Listen()
}
