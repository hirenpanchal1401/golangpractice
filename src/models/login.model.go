package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// Login ...
type Login struct {
	gorm.Model
	Username string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(100);unique_index"`
}

type Claims struct {
	Username string `gorm:"type:varchar(50)"`
	jwt.StandardClaims
}
