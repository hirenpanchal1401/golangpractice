package models

import (
	"github.com/jinzhu/gorm"
)

// Employee ...
type Employee struct {
	gorm.Model
	Name  string `gorm:"type:varchar(50)"`
	Age   int
	Email string `gorm:"type:varchar(100);unique_index"`
}
