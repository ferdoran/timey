package service

import (
	"errors"
	"fmt"
)

type InMemoryRepository[K comparable, V any] struct {
	entries map[K]*V
}

func NewInMemoryRepository[K comparable, V any]() InMemoryRepository[K, V] {
	return InMemoryRepository[K, V]{make(map[K]*V)}
}

func (r *InMemoryRepository[K, V]) GetAll() (values []V, err error) {
	values = make([]V, 0)
	for _, v := range r.entries {
		values = append(values, *v)
	}
	return
}

func (r *InMemoryRepository[K, V]) Get(key K) (value V, err error) {
	val, exists := r.entries[key]
	if !exists {
		err = errors.New(fmt.Sprintf("no entry for key %v", key))
	}
	value = *val
	return
}

func (r *InMemoryRepository[K, V]) Create(key K, value *V) (V, error) {
	if _, exists := r.entries[key]; exists {
		return *value, errors.New(fmt.Sprintf("entry for key %v already exists", key))
	}

	r.entries[key] = value
	return *value, nil
}

func (r *InMemoryRepository[K, V]) Delete(key K) error {
	_, exists := r.entries[key]
	if !exists {
		return errors.New(fmt.Sprintf("no entry for key %v", key))
	}

	delete(r.entries, key)
	return nil
}

func (r *InMemoryRepository[K, V]) Update(key K, value *V) error {
	r.entries[key] = value
	return nil
}
