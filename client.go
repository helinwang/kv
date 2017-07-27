package kv

import "net/rpc"

// Client is the client to the kv service.
type Client struct {
	c *rpc.Client
}

// New creates a new kv client.
func New(addr string) (*Client, error) {
	c, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &Client{c: c}, nil
}

// Put puts value into kv service.
func (c *Client) Put(key []byte, value []byte) error {
	kv := KV{Key: key, Value: value}
	return c.c.Call("Service.Put", kv, nil)
}

// Get gets value from kv service.
func (c *Client) Get(key []byte) ([]byte, error) {
	var value []byte
	err := c.c.Call("Service.Get", key, &value)
	if err != nil {
		return nil, err
	}

	return value, err
}
