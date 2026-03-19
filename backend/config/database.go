package config

import (

	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)

var DB *gorm.DB

func ConnectDatabase() {
	err:= godotenv.Load()
	if err != nil {
		log.Println("Belum ada file .env, pake environment variable OS")
	}	

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek ke database bro!", err)
	}

	fmt.Println("Database berhasil terhubung")
	DB = database
		
}