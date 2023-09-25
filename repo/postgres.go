package repo

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	host     = os.Getenv("HOST")
	port     = os.Getenv("DB_POSTGRES_PORT")
	user     = os.Getenv("DB_POSTGRES_USER")
	password = os.Getenv("DB_POSTGRES_PASSWORD")
	dbname   = os.Getenv("DB_POSTGRES_NAME")
)

func Conn() *gorm.DB {
	logger := logger.Default
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		panic(err)
	}
	return db
}
