package service

import (
	"errors"
	"fmt"
	"timey/model"
)

type StatementOfWorkRepository interface {
	CRUDRepository[string, model.StatementOfWork]
	GetByCustomerID(customerId string) ([]model.StatementOfWork, error)
	GetByCustomerIDAndSowID(customerId, sowId string) (model.StatementOfWork, error)
}

type InMemoryStatementOfWorkRepository struct {
	InMemoryRepository[string, model.StatementOfWork]
}
type GormStatementOfWorkRepository struct {
	GormRepository[string, model.StatementOfWork]
}

func (imr *InMemoryStatementOfWorkRepository) GetByCustomerID(customerId string) (values []model.StatementOfWork, err error) {
	values = make([]model.StatementOfWork, 0)

	for _, v := range imr.entries {
		if v.CustomerID == customerId {
			values = append(values, *v)
		}
	}

	return
}

func (imr *InMemoryStatementOfWorkRepository) GetByCustomerIDAndSowID(customerId, sowId string) (model.StatementOfWork, error) {
	for _, v := range imr.entries {
		if v.CustomerID == customerId && v.ID == sowId {
			return *v, nil
		}
	}

	return model.StatementOfWork{}, errors.New(fmt.Sprintf("no sow for customer %s with id %s", customerId, sowId))
}

func (gr *GormStatementOfWorkRepository) GetByCustomerID(customerId string) (values []model.StatementOfWork, err error) {
	values = make([]model.StatementOfWork, 0)
	err = gr.db.Where(&model.StatementOfWork{CustomerID: customerId}).Find(&values).Error

	return
}

func (gr *GormStatementOfWorkRepository) GetByCustomerIDAndSowID(customerId, sowId string) (value model.StatementOfWork, err error) {
	err = gr.db.Where(&model.StatementOfWork{ID: sowId, CustomerID: customerId}).First(&value).Error
	return
}
