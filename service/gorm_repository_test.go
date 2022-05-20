package service

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"testing"
)

type TestTable struct {
	ID        uint   `gorm:"primaryKey,autoIncrement"`
	SomeValue string `gorm:"not null"`
}

var db *gorm.DB

func TestMain(m *testing.M) {
	var err error
	db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	code := m.Run()
	os.Exit(code)
}

func TestNewGormRepository(t *testing.T) {
	t.Run("initialises empty repository", func(t *testing.T) {
		repo := NewGormRepository[uint, TestTable](db, TestTable{})

		values, err := repo.GetAll()
		if err != nil {
			t.Error(err)
		}

		if size := len(values); size != 0 {
			t.Errorf("expected repository to have 0 values but got %d", size)
		}
	})
}

func TestGormRepository_Create(t *testing.T) {
	repo := NewGormRepository[uint, TestTable](db, TestTable{})

	t.Run("creates a value in the repository with its id", func(t *testing.T) {
		val := TestTable{SomeValue: "test"}
		if created, err := repo.Create(0, &val); err != nil {
			t.Error(err)
		} else {
			val = created
		}

		var values []TestTable
		result := db.Find(&values)

		if result.Error != nil {
			t.Error(result.Error)
		}

		defer db.Delete(&values)

		if result.RowsAffected != 1 || len(values) != 1 {
			t.Errorf("expected 1 row but got %d", result.RowsAffected)
		}

		actualValue := values[0]
		if actualValue.ID == 0 {
			t.Errorf("ID was not set properly. Got %d", actualValue.ID)
		}

		if actualValue.SomeValue != "test" {
			t.Errorf("value was not set properly: %s", actualValue.SomeValue)
		}
	})
}

func TestGormRepository_GetAll(t *testing.T) {
	repo := NewGormRepository[uint, TestTable](db, TestTable{})
	expected := TestTable{SomeValue: "test"}
	if insertedVal, err := repo.Create(0, &expected); err != nil {
		t.Error(err)
	} else {
		expected = insertedVal
	}
	defer db.Delete(&expected)

	t.Run("creates a value in the repository with its id", func(t *testing.T) {
		result, err := repo.GetAll()

		if err != nil {
			t.Error(err)
		}

		if size := len(result); size != 1 {
			t.Errorf("expected 1 row but got %d", size)
		}

		actualValue := result[0]
		if actualValue.ID != expected.ID {
			t.Errorf("ID was not set properly. Got %d", actualValue.ID)
		}

		if actualValue.SomeValue != "test" {
			t.Errorf("value was not set properly: %s", actualValue.SomeValue)
		}
	})
}

func TestGormRepository_Get(t *testing.T) {
	repo := NewGormRepository[uint, TestTable](db, TestTable{})
	expected := TestTable{SomeValue: "test"}
	if insertedVal, err := repo.Create(0, &expected); err != nil {
		t.Error(err)
	} else {
		expected = insertedVal
	}
	defer db.Delete(&expected)

	t.Run("creates a value in the repository with its id", func(t *testing.T) {
		result, err := repo.Get(expected.ID)

		if err != nil {
			t.Error(err)
		}

		if result != expected {
			t.Errorf("values did not match. expected %v, but got %v", expected, result)
		}
	})
}

func TestGormRepository_Update(t *testing.T) {
	repo := NewGormRepository[uint, TestTable](db, TestTable{})
	initial := TestTable{SomeValue: "test"}
	if insertedVal, err := repo.Create(0, &initial); err != nil {
		t.Error(err)
	} else {
		initial = insertedVal
	}
	defer db.Delete(&initial)

	t.Run("updates a value in the repository", func(t *testing.T) {
		update := TestTable{ID: initial.ID, SomeValue: "updated"}
		err := repo.Update(update.ID, &update)
		if err != nil {
			t.Error(err)
		}

		result, err := repo.Get(update.ID)
		if err != nil {
			t.Error(err)
		}

		if result != update {
			t.Errorf("values did not match. initial %v, but got %v", update, result)
		}
	})
}

func TestGormRepository_Delete(t *testing.T) {
	repo := NewGormRepository[uint, TestTable](db, TestTable{})
	initial := TestTable{SomeValue: "test"}
	if insertedVal, err := repo.Create(0, &initial); err != nil {
		t.Error(err)
	} else {
		initial = insertedVal
	}

	t.Run("deletes an entry in the repository", func(t *testing.T) {
		err := repo.Delete(initial.ID)
		if err != nil {
			t.Error(err)
		}

		result, err := repo.Get(initial.ID)
		if err == nil {
			t.Errorf("expecting value to be already deleted: %v", result)
		}
	})
}
