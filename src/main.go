package main

import (
	"CRUD_GORM/src/migrations"
	"CRUD_GORM/src/routes"
	"log"
	"net/http"
)

func main() {
	migrations.Migrate()
	routes.Route().HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	})

	log.Println("Server running on port : 8080")
	http.ListenAndServe(":8080", routes.Route())
}