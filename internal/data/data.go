package data

import (
	"errors"
	"sync"
)

var (
	ErrEmptyID   = errors.New("id cannot be empty")
	ErrEmptyItem = errors.New("item cannot be empty")
)

var mutex = &sync.RWMutex{}

type storage struct {
	items map[string]interface{}
}

func NewInMemoryStorage() storage {
	return storage{
		items: make(map[string]interface{}),
	}
}

func (s storage) Set(id string, item interface{}) error {
	mutex.Lock()
	defer mutex.Unlock()

	if id == "" {
		return ErrEmptyID
	}
	if item == nil {
		return ErrEmptyItem
	}
	s.items[id] = item
	return nil
}

func (s storage) Get(id string) (interface{}, bool) {
	mutex.RLock()
	defer mutex.RUnlock()

	item, found := s.items[id]
	return item, found
}

func (s storage) Delete(id string) bool {
	mutex.Lock()
	defer mutex.Unlock()

	_, found := s.items[id]
	if found {
		delete(s.items, id)
	}
	return found
}
