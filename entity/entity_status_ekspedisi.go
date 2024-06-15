package entity

import "time"

type StatusEkspedisi struct {
	ID        int
	ResiInfo  string
	UserID    int
	User      User
	OrderID   int
	OngkirID  int
	Order     Order                 `gorm:"foreignKey:OrderID"`
	Ongkir    ApplyShippingResponse `gorm:"foreignKey:OngkirID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
