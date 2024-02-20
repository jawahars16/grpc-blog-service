package data_test

import (
	"testing"

	"github.com/jawahars16/grpc-blog-service/internal/data"
	"github.com/stretchr/testify/assert"
)

func Test_storage_Set(t *testing.T) {
	t.Run("given an item to set function, it should store the item in the storage", func(t *testing.T) {
		storage := data.NewInMemoryStorage()
		err := storage.Set("1", "item1")
		assert.NoError(t, err)
	})

	t.Run("given an item with an existing id, it should update the item in the storage", func(t *testing.T) {
		storage := data.NewInMemoryStorage()
		err := storage.Set("1", "item1")
		assert.NoError(t, err)
		err = storage.Set("1", "item2")
		assert.NoError(t, err)
	})

	t.Run("given an item with an empty id, it should return an error", func(t *testing.T) {
		storage := data.NewInMemoryStorage()
		err := storage.Set("", "item1")
		assert.ErrorIs(t, err, data.ErrEmptyID)
	})

	t.Run("given an item with an empty value, it should return an error", func(t *testing.T) {
		storage := data.NewInMemoryStorage()
		err := storage.Set("1", nil)
		assert.ErrorIs(t, err, data.ErrEmptyItem)
	})

	t.Run("given an ID to delete function, it should remove the item from the storage", func(t *testing.T) {
		storage := data.NewInMemoryStorage()
		err := storage.Set("1", "item1")
		assert.NoError(t, err)
		deleted := storage.Delete("1")
		assert.True(t, deleted)
		_, found := storage.Get("1")
		assert.False(t, found)
	})

	t.Run("given an ID to delete that does not exist in the storage, it should return false", func(t *testing.T) {
		storage := data.NewInMemoryStorage()
		deleted := storage.Delete("1")
		assert.False(t, deleted)
	})

	t.Run("given an ID to get function, it should return the item from the storage", func(t *testing.T) {
		storage := data.NewInMemoryStorage()
		err := storage.Set("1", "item1")
		assert.NoError(t, err)
		item, found := storage.Get("1")
		assert.True(t, found)
		assert.Equal(t, "item1", item)
	})

	t.Run("given an ID to get that does not exist in the storage, it should return false", func(t *testing.T) {
		storage := data.NewInMemoryStorage()
		_, found := storage.Get("1")
		assert.False(t, found)
	})
}
