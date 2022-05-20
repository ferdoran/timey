package service

type CRUDRepository[K comparable, V any] interface {
	GetAll() (values []V, err error)
	Get(key K) (value V, err error)
	Create(key K, value *V) (V, error)
	Update(key K, value *V) error
	Delete(key K) error
}
