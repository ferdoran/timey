package service

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GormRepository[K comparable, V any] struct {
	db       *gorm.DB
	emptyVal V
}

func NewGormRepository[K comparable, V any](db *gorm.DB, emptyVal V) GormRepository[K, V] {
	repo := GormRepository[K, V]{db, emptyVal}
	if err := db.AutoMigrate(emptyVal); err != nil {
		logrus.Panic(err)
	}
	return repo
}

func (r *GormRepository[K, V]) GetAll() (values []V, err error) {
	if err = r.db.Find(&values).Error; err != nil {
		return nil, err
	}
	return values, nil
}

func (r *GormRepository[K, V]) Get(key K) (value V, err error) {
	if err = r.db.First(&value, key).Error; err != nil {
		return r.emptyVal, err
	}

	return value, nil
}

func (r *GormRepository[K, V]) Create(_ K, value *V) (V, error) {
	err := r.db.Create(value).Error
	return *value, err
}

func (r *GormRepository[K, V]) Update(_ K, value *V) error {
	return r.db.Save(value).Error
}

func (r *GormRepository[K, V]) Delete(key K) error {
	return r.db.Delete(r.emptyVal, key).Error
}
