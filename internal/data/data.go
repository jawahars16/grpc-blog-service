package data

type storage struct {
	items map[string]interface{}
}

func NewInMemoryStorage() storage {
	return storage{
		items: make(map[string]interface{}),
	}
}

func (s storage) Set(id string, item interface{}) error {
	return nil
}

func (s storage) Get(id string) (interface{}, bool) {
	return nil, false
}

func (s storage) Delete(id string) bool {
	return false
}
