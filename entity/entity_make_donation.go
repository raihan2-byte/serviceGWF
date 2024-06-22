package entity

import "time"

type MakeDonation struct {
	ID        string `gorm:"type:varchar(36);primaryKey"`
	UserID    int
	Name      string
	Amount    int
	Message   string
	User      User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
