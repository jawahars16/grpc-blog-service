package data

import "errors"

var (
	ErrEmptyID   = errors.New("id cannot be empty")
	ErrEmptyItem = errors.New("item cannot be empty")
)

type storage struct {
	items map[string]interface{}
}

func NewInMemoryStorage() storage {
	return storage{
		items: make(map[string]interface{}),
	}
}

func (s storage) Set(id string, item interface{}) error {
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
	item, found := s.items[id]
	return item, found
}

func (s storage) Delete(id string) bool {
	return false
}
