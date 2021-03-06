package middleware

import (
	"CRUD_GORM/src/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func Auth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		err1 := godotenv.Load()
		if err1 != nil {
			log.Fatal("Error loading .env file")
		}
		var jwtKey = []byte(os.Getenv("JWT_KEY"))
		// We can obtain the session token from the requests cookies, which come with every request
		c, err := req.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				res.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		// Get the JWT string from the cookie
		tknStr := c.Value
		// Initialize a new instance of `Claims`
		claims := models.Claims{}
		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(tknStr, &claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if !tkn.Valid {
			res.WriteHeader(http.StatusUnauthorized)
			return
		}
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				res.WriteHeader(http.StatusUnauthorized)
				return
			}
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		next.ServeHTTP(res, req)
	})

	// next.ServeHTTP(res, req)
	// Finally, return the welcome message to the user, along with their
	// username given in the token
	// res.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}
