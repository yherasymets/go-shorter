package repo

import (
	"fmt"
	"log"
	"os"

	"github.com/yherasymets/go-shorter/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Retrieve PostgreSQL connection parameters from environment variables.
var (
	host     = os.Getenv("HOST")                 // PostgreSQL server hostname or IP address.
	port     = os.Getenv("DB_POSTGRES_PORT")     // Port number for PostgreSQL server.
	user     = os.Getenv("DB_POSTGRES_USER")     // Username for PostgreSQL authentication.
	password = os.Getenv("DB_POSTGRES_PASSWORD") // Password for PostgreSQL authentication.
	dbname   = os.Getenv("DB_POSTGRES_NAME")     // Name of the PostgreSQL database.
)

// NewRepository initializes a new database connection using GORM and performs auto-migration.
func NewRepository() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(service.Link{}); err != nil {
		log.Printf("migration failed: %v", err)
	}
	return db
}
