package cache

import (
	"log"
	"sync"
)

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

// 创建Cache接口的一个实现结构体实例(inMemoryCache)
func newInMemoryCache() *inMemoryCache {
	return &inMemoryCache{make(map[string][]byte), sync.RWMutex{}, Stat{}}
}

// 缓存接口
type Cache interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
	GetStat() Stat
}

// 缓存接口实现结构体
type inMemoryCache struct {
	container map[string][]byte
	mutex     sync.RWMutex
	Stat
}

// inMemoryCache实现设置缓存方法
func (i *inMemoryCache) Set(key string, value []byte) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	tmp, exist := i.container[key]
	if exist {
		i.del(key, tmp)
	}
	i.container[key] = value
	i.add(key, value)
	return nil
}

// inMemoryCache实现获取缓存方法
func (i *inMemoryCache) Get(key string) ([]byte, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	return i.container[key], nil
}

// inMemoryCache实现删除缓存方法
func (i *inMemoryCache) Del(key string) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	value, exist := i.container[key]
	if exist {
		delete(i.container, key)
		i.del(key, value)
	}
	return nil
}

// inMemoryCache实现获取状态方法
func (i *inMemoryCache) GetStat() Stat {
	return i.Stat
}

// 状态结构体
type Stat struct {
	Count     int64
	KeySize   int64
	ValueSize int64
}

func (s *Stat) add(key string, value []byte) {
	s.Count++
	s.KeySize += int64(len(key))
	s.ValueSize += int64(len(value))
}

func (s *Stat) del(key string, value []byte) {
	s.Count--
	s.KeySize -= int64(len(key))
	s.ValueSize -= int64(len(value))
}
