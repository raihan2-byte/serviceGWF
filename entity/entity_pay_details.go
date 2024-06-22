package entity

import "time"

type PaymentDetails struct {
	ID                int `gorm:"primaryKey"`
	PaymentID         int `gorm:"index"` // Add index for performance
	PaymentDonationID int `gorm:"index"` // Add index for performance
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Payment           *Payment         `gorm:"foreignKey:PaymentID"`
	PaymentDonation   *PaymentDonation `gorm:"foreignKey:PaymentDonationID"`
}
