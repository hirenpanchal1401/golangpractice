package migrations

import (
	"CRUD_GORM/src/db"
	"CRUD_GORM/src/models"
)

// Migrate ...
func Migrate() {
	db := db.DbConn()
	employee := models.Employee{}
	login := models.Login{}
	db.AutoMigrate(&employee, &login)
	// defer db.Close()
}
