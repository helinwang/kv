# kv
[![Build Status](https://travis-ci.org/helinwang/kv.svg?branch=master)](https://travis-ci.org/helinwang/kv)

A Go key/value store service based on [BoltDB](https://github.com/boltdb/bolt), access BoltDB from multiple processes.

## Integrate with Go

Please see [here](./example_test.go).

## CLI example

```bash
go get github.com/helinwang/kv/cmd/kvctl
go get github.com/helinwang/kv/cmd/kv

# Start kv service:
kv -path db.bin

# Test (open another terminal)
kvctl put :8080 hello hi
kvctl get :8080 hello

# Output: hi
```

## Graceful Shutdown

Supported in the CLI. Please see [here](./cmd/kv/main.go).
