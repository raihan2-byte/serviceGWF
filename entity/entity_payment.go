package entity

import "time"

type Payment struct {
	ID             int
	StatusPayment  string
	TransactionID  string
	OrderID        string
	UserID         int
	User           User
	Order          Order
	PaymentDetails *PaymentDetails
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
