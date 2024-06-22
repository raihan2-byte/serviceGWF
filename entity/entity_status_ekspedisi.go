package entity

import "time"

type StatusEkspedisi struct {
	ID        int
	ResiInfo  string
	UserID    int
	User      User
	OrderID   string
	OngkirID  int
	PaymentID int
	Order     Order                 `gorm:"foreignKey:OrderID"`
	Ongkir    ApplyShippingResponse `gorm:"foreignKey:OngkirID"`
	Payment   Payment
	CreatedAt time.Time
	UpdatedAt time.Time
}
