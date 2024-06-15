package database

import (
	"log"
	"payment-gwf/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {

	dsn := "root:@tcp(127.0.0.1:3306)/payment-gwf?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	// Auto Migration

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Products{})
	db.AutoMigrate(&entity.Cart{})
	db.AutoMigrate(&entity.BuyerAddress{})
	db.AutoMigrate(&entity.Order{})
	db.AutoMigrate(&entity.OrderItem{})
	db.AutoMigrate(&entity.ApplyShippingResponse{})
	db.AutoMigrate(&entity.MakeDonation{})
	db.AutoMigrate(&entity.StatusEkspedisi{})
	db.AutoMigrate(&entity.Payment{})
	// db.AutoMigrate(&entity.Transaction{})

	return db, nil

}
