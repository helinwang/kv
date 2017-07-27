package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/helinwang/kv"
)

func main() {
	path := flag.String("path", "", "path to db file, will be created if not exists")
	port := flag.Int("port", 8080, "service port")
	ip := flag.String("ip", "0.0.0.0", "service ip")
	flag.Parse()

	db, err := bolt.Open(*path, 0666, nil)
	if err != nil {
		panic(err)
	}

	s := &kv.Service{DB: db}
	rpc.Register(s)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", *ip+":"+strconv.Itoa(*port))
	if e != nil {
		log.Fatal("listen error:", e)
	}

	err = http.Serve(l, nil)
	if err != nil {
		panic(err)
	}
}
