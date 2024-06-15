package entity

import (
	"time"
)

type Order struct {
	ID          int
	UserID      int
	TotalPrice  int
	OngkirID    int
	ShippingFee int                   `json:"shipping_fee"` // Tambahkan field ini
	Items       []OrderItem           `gorm:"foreignKey:OrderID"`
	Ongkir      ApplyShippingResponse `gorm:"foreignKey:OngkirID"`
	User        User                  `gorm:"foreignKey:UserID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type OrderItem struct {
	ID        int
	OrderID   int
	ProductID int
	Quantity  int
	Price     int
	Order     Order    `gorm:"foreignKey:OrderID"`
	Product   Products `gorm:"foreignKey:ProductID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
