package context

import "testing"

func TestGet_NoBinding(t *testing.T) {
	t.Run("given no binding it should return an error", func(t *testing.T) {
		bar, err := Get[string]("foo")
		if err == nil {
			t.Error("expecting error for non existent binding, but did not receive any")
		}

		if bar != nil {
			t.Errorf("expected no binding but got one: %s", *bar)
		}
	})
}

func TestBind(t *testing.T) {
	t.Run("given no binding, it should create one", func(t *testing.T) {
		bar := "bar"
		Bind[string]("foo", &bar)

		result, err := Get[string]("foo")
		if err != nil {
			t.Error(err)
		}

		if result != &bar {
			t.Errorf("expected to get bound reference %v but got %v", &bar, result)
		}
	})
}

func TestBind_panic(t *testing.T) {
	t.Run("given an existing binding, binding another value to it panics", func(t *testing.T) {
		bar := "bar"
		Bind[string]("foo2", &bar)

		defer func() {
			if r := recover(); r == nil {
				t.Error("expected Bind to panic, but it did not")
			}
		}()

		Bind[string]("foo2", &bar)
	})
}
