package main

import (
	"github.com/thewalkers/inmemory-cache/http-server/server/cache"
	"github.com/thewalkers/inmemory-cache/http-server/server/http"
	"github.com/thewalkers/inmemory-cache/tcp-server/tcp"
)

func main() {
	ca := cache.New("inmemory")
	go tcp.New(ca).Listen()
	http.New(ca).Listen()
}
