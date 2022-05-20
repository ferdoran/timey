package model

type Customer struct {
	ID               string            `gorm:"primaryKey" json:"id" param:"id"`
	Name             string            `json:"name"`
	StatementsOfWork []StatementOfWork `gorm:"foreignKey:CustomerID" json:"sows,omitempty"`
}
