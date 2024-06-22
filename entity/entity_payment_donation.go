package entity

import "time"

type PaymentDonation struct {
	ID             int
	StatusPayment  string
	TransactionID  string
	MakeDonationID string
	MakeDonation   MakeDonation `gorm:"foreignKey:MakeDonationID"`
	UserID         int
	User           User
	PaymentDetails *PaymentDetails
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
