package service

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"timey/context"
	"timey/model"
)

const SowRepoQualifier = "sowRepository"

func InitSOWRepo() {
	logrus.Info("initialising sow repo")
	var inMemRepo = NewInMemoryRepository[string, model.StatementOfWork]()
	var repo StatementOfWorkRepository = &InMemoryStatementOfWorkRepository{inMemRepo}

	db, err := context.Get[gorm.DB]("db")
	if err == nil {
		gormRepo := NewGormRepository[string, model.StatementOfWork](db, model.StatementOfWork{})
		repo = &GormStatementOfWorkRepository{gormRepo}
	}

	logrus.Error(err)
	logrus.Warnf("using in-memory StatementOfWork repositoy")
	context.Bind(SowRepoQualifier, &repo)
}
