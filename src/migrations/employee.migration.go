package migrations

import (
	"CRUD_GORM/src/db"
	"CRUD_GORM/src/models"
)

// Migrate ...
func Migrate() {
	db := db.DbConn()
	employee := models.Employee{}
	db.AutoMigrate(&employee)
	// defer db.Close()
}
