package boltcache

import (
	"os"

	"github.com/boltdb/bolt"
)

// BoltDB template caching

type BoltCache struct {
	source *bolt.DB
}

func New(path string, mode os.FileMode, options *bolt.Options) (*BoltCache, error) {
	boltC, err := bolt.Open(path, mode, options)
	if err != nil {
		return nil, err
	}
	return &BoltCache{boltC}, nil
}

func (b *BoltCache) Get(id string) []byte {
	var val []byte
	b.source.View(func(tx *bolt.Tx) error {
		val = tx.Bucket([]byte("default")).Get([]byte(id))
		return nil
	})
	return val
}

func (b *BoltCache) Set(id string, data []byte) error {
	b.source.Update(func(tx *bolt.Tx) error {
		buck, err := tx.CreateBucketIfNotExists([]byte("default"))
		if err != nil {
			return err
		}
		buck.Put([]byte(id), data)
		return nil
	})
	return nil
}

func (b *BoltCache) Update(id string, data []byte) error {
	b.source.Update(func(tx *bolt.Tx) error {
		buck, err := tx.CreateBucketIfNotExists([]byte("default"))
		if err != nil {
			return err
		}
		buck.Put([]byte(id), data)
		return nil
	})
	return nil
}

func (b *BoltCache) Del(id string) error {
	b.source.Update(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte("default"))
		buck.Delete([]byte(id))
		return nil
	})
	return nil
}
