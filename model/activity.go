package model

import "time"

type Activity struct {
	ID          string `gorm:"primaryKey;type=uuid"`
	Ticket      string
	Description string `gorm:"not null"`
	SOW         string
	Begin       time.Time
	End         time.Time
}
