package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

var (
	db   *gorm.DB
	err  error
	user string
)

// ConnectDB ...
func ConnectDB() error {
	db, err = gorm.Open(
		"postgres",
		"host="+os.Getenv("PG_HOST")+" user="+os.Getenv("PG_USER")+" dbname="+os.Getenv("PG_DBNAME")+" sslmode=disable password="+os.Getenv("PG_PASSWORD")+"")
	if err != nil {
		return err
	}

	db.LogMode(true)
	log.Println("Connection Established")
	return nil
}

// GetDB ...
func GetDB() *gorm.DB {
	if db == nil {
		log.Panicln(db)
	}
	return db
}

// CloseDB ...
func CloseDB() {
	db.Close()
}
