package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/MafuSora/portofolio_db/helper"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env File")
	}

	host := os.Getenv("HOST_DB")
	init_port := os.Getenv("PORT_DB")
	port, err2 := strconv.Atoi(init_port)
	if err2 != nil {
		log.Fatal("Error loading .env File")
	}
	user := os.Getenv("USERNAME_DB")
	password := os.Getenv("PASSWORD_DB")
	dbname := os.Getenv("DB_NAME_DB")

	sqlInfo := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", user, password, host, port, dbname)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)
	return db
}
