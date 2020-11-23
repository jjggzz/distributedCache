package main

import (
	"Fool/main/cache"
	"Fool/main/http"
)

func main() {
	http.New(cache.New("inMemory")).Listen()
}
