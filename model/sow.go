package model

type StatementOfWork struct {
	ID         string     `gorm:"primaryKey" json:"id" param:"sowId"`
	Name       string     `json:"name"`
	CustomerID string     `json:"customerId"`
	Activities []Activity `gorm:"foreignKey:SOW" json:"activities,omitempty"`
}
