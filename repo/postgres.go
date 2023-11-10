package repo

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yherasymets/go-shorter/internal/shorter"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("HOST")
	port     = os.Getenv("DB_POSTGRES_PORT")
	user     = os.Getenv("DB_POSTGRES_USER")
	password = os.Getenv("DB_POSTGRES_PASSWORD")
	dbname   = os.Getenv("DB_POSTGRES_NAME")
)

func Connection() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(shorter.Link{}); err != nil {
		logrus.Errorf("migration failed: %v", err)
	}
	return db
}
