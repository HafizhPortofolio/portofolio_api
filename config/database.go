package config

import (
	"fmt"

	"github.com/m/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "db.kjmsxmqosydoptrignaj.supabase.co"
	port     = 5432
	user     = "postgres"
	password = "@orang111357"
	dbname   = "postgres"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", user, password, host, port, dbname)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)
	return db
}
