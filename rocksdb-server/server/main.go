package main

import (
	"flag"
	"log"

	"github.com/thewalkers/inmemory-cache/http-server/server/http"
	"github.com/thewalkers/inmemory-cache/rocksdb-server/server/cache"
	"github.com/thewalkers/inmemory-cache/tcp-server/tcp"
)

func main() {
	typ := flag.String("type", "inmemory", "cache type")
	flag.Parse()
	log.Println("type is", *typ)
	c := cache.New(*typ)
	go tcp.New(c).Listen()
	http.New(c).Listen()
}
