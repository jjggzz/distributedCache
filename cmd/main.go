package main

import (
	"distributeCache/cache"
	"distributeCache/http"
)

func main() {
	http.New(cache.New("inMemory")).Listen()
}
