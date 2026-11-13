package config

import "os"

type HTTP struct {
	Addr string
}

func loadHTTP() HTTP {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}
	return HTTP{Addr: addr}
}
