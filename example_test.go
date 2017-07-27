package kv_test

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"github.com/boltdb/bolt"
	"github.com/helinwang/kv"
)

func Example() {
	go func() {
		// Start the service.
		db, err := bolt.Open("/tmp/db_test.bin", 0666, nil)
		if err != nil {
			panic(err)
		}

		s := &kv.Service{DB: db}
		rpc.Register(s)
		rpc.HandleHTTP()

		l, e := net.Listen("tcp", ":8081")
		if e != nil {
			log.Fatal("listen error:", e)
		}

		err = http.Serve(l, nil)
		if err != nil {
			panic(err)
		}
	}()

	// Wait for the service to start.
	time.Sleep(50 * time.Millisecond)

	c, err := kv.New(":8081")
	if err != nil {
		panic(err)
	}

	err = c.Put([]byte("hello"), []byte("hi"))
	if err != nil {
		panic(err)
	}

	v, err := c.Get([]byte("hello"))
	if err != nil {
		panic(err)
	}

	fmt.Println(string(v))

	// Output:
	// hi
}
