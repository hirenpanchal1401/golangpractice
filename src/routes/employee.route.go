package routes

import (
	"CRUD_GORM/src/controllers"
	"github.com/gorilla/mux"
)

//Route ...
func Route() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.GetAllEmployee).Methods("GET")
	router.HandleFunc("/create", controller.CreateEmployee).Methods("POST")
	router.HandleFunc("/findbyname", controller.GetEmployee).Methods("GET")
	router.HandleFunc("/updatebyid/{id}", controller.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/deletebyid/{id}", controller.DeleteEmployee).Methods("DELETE")
	return router
}
