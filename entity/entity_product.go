package entity

import "time"

type Products struct {
	ID        int
	Name      string
	Price     int
	Stock     int
	FileName  []ProductImage `gorm:"foreignKey:ProductID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductImage struct {
	ID        int `gorm:"primaryKey"`
	FileName  string
	ProductID int
	CreatedAt time.Time
	UpdatedAt time.Time
}
