package main

import (
	"github.com/thewalkers/inmemory-cache/http-server/server/cache"
	"github.com/thewalkers/inmemory-cache/http-server/server/http"
)

func main() {
	c := cache.New("inmemory")
	http.New(c).Listen()
}
