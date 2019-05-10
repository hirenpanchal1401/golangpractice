package routes

import (
	"CRUD_GORM/src/controllers"
	"CRUD_GORM/src/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

//Route ...
func Route() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/", Middleware(http.HandlerFunc(controller.GetAllEmployee), middleware.Auth)).Methods("GET")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.Handle("/create", Middleware(http.HandlerFunc(controller.CreateEmployee), middleware.Auth)).Methods("POST")
	router.Handle("/findbyname", Middleware(http.HandlerFunc(controller.GetEmployee), middleware.Auth)).Methods("GET")
	router.Handle("/updatebyid/{id}", Middleware(http.HandlerFunc(controller.UpdateEmployee), middleware.Auth)).Methods("PUT")
	router.Handle("/deletebyid/{id}", Middleware(http.HandlerFunc(controller.DeleteEmployee), middleware.Auth)).Methods("DELETE")
	return router
}

func Middleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}
