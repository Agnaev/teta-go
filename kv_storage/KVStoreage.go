package kv_storage

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
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
	val, ok := storage.items[key]
	if !ok {
		return nil, errors.New("No such key")
	}
	return val, nil
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

func Run() {
	ctx := context.Background()
	storage := Storage{
		make(map[string]interface{}),
		sync.RWMutex{},
	}
	go func() {
		for i := 0; i < 100; i++ {
			go storage.Put(ctx, fmt.Sprint(i), i)
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			go func(i int) {
				val, err := storage.Get(ctx, fmt.Sprint(i))
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(val)
			}(i)
		}
	}()
	time.Sleep(time.Second * 3)
}
