package entity

import "time"

type MakeDonation struct {
	ID            int
	UserID        int
	Name          string
	Amount        int
	Message       string
	StatusPayment string
	User          User `gorm:"foreignKey:UserID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
