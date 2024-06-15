package entity

import "time"

type Cart struct {
	ID         int
	UserID     int
	ProductID  int
	Quantity   int
	TotalPrice int
	Product    Products `gorm:"foreignKey:ProductID"`
	User       User     `gorm:"foreignKey:UserID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
