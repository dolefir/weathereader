package models

import (
	"github.com/simplewayUA/weathereader/db"
)

// Migrate AutoMigrate
func Migrate() {
	var getDB = db.GetDB()
	getDB.AutoMigrate(&User{})
	getDB.AutoMigrate(&UserCity{})
	getDB.AutoMigrate(&Weather{})
}
