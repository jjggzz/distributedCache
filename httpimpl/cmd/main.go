package main

import (
	"distributeCache/httpimpl/cache"
	http2 "distributeCache/httpimpl/http"
)

func main() {
	http2.New(cache.New("inMemory")).Listen()
}
