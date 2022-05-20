package service

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"timey/context"
	"timey/model"
)

const CustomerRepoQualifier = "customerRepository"

func InitCustomerRepo() {
	logrus.Info("initialising customer repo")
	var inMemRepo = NewInMemoryRepository[string, model.Customer]()
	var repo CRUDRepository[string, model.Customer] = &inMemRepo

	db, err := context.Get[gorm.DB]("db")
	if err == nil {
		gormRepo := NewGormRepository[string, model.Customer](db, model.Customer{})
		repo = &gormRepo
	}

	logrus.Error(err)
	logrus.Warnf("using in-memory Customer repositoy")

	context.Bind[CRUDRepository[string, model.Customer]](CustomerRepoQualifier, &repo)
}
