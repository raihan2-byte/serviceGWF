package database

import (
	"log"
	"os"
	"payment-gwf/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {

	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); exists == false {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file:", err)
		}
	}

	databaseURL := os.Getenv("DATABASE_URL")

	// Check if DATABASE_URL is set
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// Open the database connection using the URL
	db, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error:", err)
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
	db.AutoMigrate(&entity.DoPayment{})
	// db.AutoMigrate(&entity.VaNumber{})
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
