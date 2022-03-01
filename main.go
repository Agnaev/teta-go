package main

import (
	"flag"
	"test/ftp/client"
	"test/ftp/server"
	"test/kv_storage"
)

func main() {
	val := flag.String("service", "", "")
	flag.Parse()
	switch *val {
	case "client":
		client.Run()
	case "server":
		server.Run()
	case "kv-store":
		kv_storage.Run()
	}

}
