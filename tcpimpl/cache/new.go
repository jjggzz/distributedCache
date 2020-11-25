package cache

import "log"

func New(typ string) Cache {
	var c Cache
	if typ == "inMemory" {
		c = newInMemoryCache()
	}
	if c == nil {
		panic("未知的缓存类型" + typ)
	}
	log.Print(typ, "读取服务")

	return c
}
