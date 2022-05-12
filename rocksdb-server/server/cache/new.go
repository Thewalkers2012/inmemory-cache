package cache

import (
	"log"

	"github.com/thewalkers/inmemory-cache/http-server/server/cache"
)

func New(typ string) cache.Cache {
	var c cache.Cache
	if typ == "rocksdb" {
		c = newRocksdbCache()
	}
	if c == nil {
		panic("unknown cache type " + typ)
	}
	log.Println(typ, "ready to serve")
	return c
}
