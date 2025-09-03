package main

import (
	"github.com/joho/godotenv"
	"go/adv-dev/internal/link"
	"go/adv-dev/internal/stat"
	"go/adv-dev/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&link.Link{}, &user.User{}, &stat.Stat{})
	if err != nil {
		panic(err)
	}
}
