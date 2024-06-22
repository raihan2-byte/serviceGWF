package database

import (
	"log"
	"payment-gwf/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {

	dsn := "root:@tcp(127.0.0.1:3306)/pay?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	// Auto Migration

	// db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Products{})
	db.AutoMigrate(&entity.Cart{})
	db.AutoMigrate(&entity.BuyerAddress{})
	db.AutoMigrate(&entity.Order{})
	db.AutoMigrate(&entity.OrderItem{})
	db.AutoMigrate(&entity.ApplyShippingResponse{})
	errs := db.AutoMigrate(&entity.MakeDonation{})
	if errs != nil {
		// Tangani kesalahan di sini, misalnya dengan mencetak pesan kesalahan atau mengembalikan kesalahan
		log.Fatalf("Error during migration: %v", errs)
	}

	db.AutoMigrate(&entity.StatusEkspedisi{})
	db.AutoMigrate(&entity.Payment{})
	db.AutoMigrate(&entity.ProductImage{})
	errs = db.AutoMigrate(&entity.PaymentDonation{})
	if errs != nil {
		// Tangani kesalahan di sini, misalnya dengan mencetak pesan kesalahan atau mengembalikan kesalahan
		log.Fatalf("Error during migration: %v", errs)
	}
	db.AutoMigrate(&entity.PaymentDetails{})
	// db.AutoMigrate(&entity.Transaction{})

	return db, nil

}
