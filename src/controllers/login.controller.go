package controller

import (
	"CRUD_GORM/src/db"
	"CRUD_GORM/src/models"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

//Login ...
func Login(res http.ResponseWriter, req *http.Request) {

	db := db.DbConn()
	user := models.Login{}
	req.ParseForm()
	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	userlogged := db.Where("username = ? AND password = ?", &user.Username, &user.Password).Find(&user)
	if userlogged.Error != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := models.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Loading .env file
	err1 := godotenv.Load()
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}

	var jwtKey = []byte(os.Getenv("JWT_KEY"))

	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	type TokenString struct {
		Token string
	}

	// json.NewEncoder(res).Encode(TokenString{Token: tokenString})

	http.SetCookie(res, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	defer db.Close()
}
