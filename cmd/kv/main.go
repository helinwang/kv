package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"strconv"
	"time"

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
	rpcSrv := rpc.NewServer()
	err = rpcSrv.Register(s)
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    *ip + ":" + strconv.Itoa(*port),
		Handler: rpcSrv,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal("listen error:", err)
		}
	}()

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	cancel()
}
