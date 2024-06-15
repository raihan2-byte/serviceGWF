package entity

import "time"

type Payment struct {
	ID            int
	StatusPayment string
	OrderID       int
	UserID        int
	User          User
	Order         Order
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
