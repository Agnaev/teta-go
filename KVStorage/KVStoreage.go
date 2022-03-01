package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

type KVStorage interface {
	Get(context.Context, string) (interface{}, error)
	Put(context.Context, string, interface{}) error
	Delete(context.Context, string) error
}

type Storage struct { // implements KVStorage
	items map[string]interface{}
	lock  sync.RWMutex
}

func (storage *Storage) Get(ctx context.Context, key string) (interface{}, error) {
	storage.lock.RLock()
	defer storage.lock.RUnlock()
	if val, ok := storage.items[key]; ok {
		return val, nil
	}
	return nil, errors.New("No such key")
}

func (storage *Storage) Put(ctx context.Context, key string, val interface{}) error {
	storage.lock.Lock()
	defer storage.lock.Unlock()
	storage.items[key] = val
	return nil
}

func (storage *Storage) Delete(ctx context.Context, key string) error {
	storage.lock.Lock()
	defer storage.lock.Unlock()
	if _, ok := storage.items[key]; !ok {
		return errors.New("Map does not contains key")
	}
	delete(storage.items, key)
	return nil
}

func main() {
	ctx := context.Background()
	storage := Storage{
		make(map[string]interface{}),
		sync.RWMutex{},
	}
	if err := storage.Put(ctx, "1", 1); err != nil {
		fmt.Printf("Put error")
	}
	if val, err := storage.Get(ctx, "1"); err != nil {
		fmt.Printf("not key in storage")
	} else {
		fmt.Printf("%d", val)
	}
}
