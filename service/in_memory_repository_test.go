package service

import (
	"fmt"
	"testing"
)

func TestNewInMemoryRepository(t *testing.T) {
	t.Run("fresh in-memory repository is empty", func(t *testing.T) {
		repo := NewInMemoryRepository[string, string]()

		if len(repo.entries) != 0 {
			t.Errorf("repository is not empty: %d", len(repo.entries))
		}
	})
}

func TestInMemoryRepository_Create(t *testing.T) {
	input := map[string]string{
		"foo":   "bar",
		"bar":   "foo",
		"hello": "world",
	}

	repo := NewInMemoryRepository[string, string]()
	for k, v := range input {
		t.Run(fmt.Sprintf("saves entry: %s -> %s", k, v), func(t *testing.T) {
			_, err := repo.Create(k, &v)
			if err != nil {
				t.Error(err)
			}

			if repo.entries[k] != &v {
				t.Errorf("expected entry %s -> %s but got %s -> %v", k, v, k, repo.entries[k])
			}
		})
	}
}

func TestInMemoryRepository_Create_FailOnDuplicate(t *testing.T) {
	input := map[string]string{
		"foo":   "bar",
		"bar":   "foo",
		"hello": "world",
	}

	repo := NewInMemoryRepository[string, string]()
	for k, v := range input {
		repo.entries[k] = &v
		t.Run(fmt.Sprintf("saves entry: %s -> %s", k, v), func(t *testing.T) {
			_, err := repo.Create(k, &v)
			if err == nil {
				t.Errorf("expecting error for already existing entry %s -> %s", k, v)
			}
		})
	}
}
