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

	dbUsername := os.Getenv("MYSQLUSER")
	dbPassword := os.Getenv("MYSQLPASSWORD")
	dbHost := os.Getenv("MYSQLHOST")
	dbPort := os.Getenv("MYSQLPORT")
	dbName := os.Getenv("MYSQLDATABASE")

	// dsn := "root:@tcp(127.0.0.1:3306)/gwf-apps?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("DB Connection Error")
	// }

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error", err)
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
