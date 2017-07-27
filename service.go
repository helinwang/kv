package kv

import (
	"github.com/boltdb/bolt"
)

// Service is the kv service.
type Service struct {
	DB *bolt.DB
}

// KV is the argument for put.
type KV struct {
	Key   []byte
	Value []byte
}

var bucket = []byte("main")

// Put puts value into the service.
func (s *Service) Put(kv KV, _ int) error {
	return s.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}

		err = b.Put(kv.Key, kv.Value)
		return err
	})
}

// Get gets value from the service.
func (s *Service) Get(key []byte, value *[]byte) error {
	return s.DB.View(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}

		*value = b.Get(key)
		return nil
	})
}
