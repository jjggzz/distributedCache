package main

import (
	"distributeCache/tcpimpl/cache"
	"distributeCache/tcpimpl/tcp"
)

func main() {
	tcp.New(cache.New("inMemory")).Listen()
}
