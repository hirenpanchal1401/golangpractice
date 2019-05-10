package controller

import (
	"CRUD_GORM/src/db"
	"CRUD_GORM/src/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

//GetAllEmployee ...
func GetAllEmployee(res http.ResponseWriter, req *http.Request) {
	db := db.DbConn()
	employee := []models.Employee{}
	employeeList := db.Find(&employee)
	json.NewEncoder(res).Encode(employeeList.Value)
	defer db.Close()
}

//CreateEmployee ...
func CreateEmployee(res http.ResponseWriter, req *http.Request) {
	db := db.DbConn()
	employee := models.Employee{}
	req.ParseForm()
	err := json.NewDecoder(req.Body).Decode(&employee)
	if err != nil {
		log.Println(err)
	}

	emp := models.Employee{Name: employee.Name, Age: employee.Age, Email: employee.Email}
	db.Create(&emp)
	log.Println("Data Inserted", employee)
	json.NewEncoder(res).Encode(emp)
	defer db.Close()
}

//GetEmployee ...
func GetEmployee(res http.ResponseWriter, req *http.Request) {
	db := db.DbConn()
	employee := []models.Employee{}
	emp := models.Employee{}
	req.ParseForm()
	err := json.NewDecoder(req.Body).Decode(&emp)
	if err != nil {
		log.Println(err)
	}
	employeeList := db.Where("name = ?", &emp.Name).Find(&employee)
	json.NewEncoder(res).Encode(employeeList.Value)
	defer db.Close()
}

//UpdateEmployee ...
func UpdateEmployee(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	db := db.DbConn()
	employee := models.Employee{}
	req.ParseForm()
	err := json.NewDecoder(req.Body).Decode(&employee)
	if err != nil {
		log.Println(err)
	}
	emp := models.Employee{Name: employee.Name, Age: employee.Age, Email: employee.Email}
	db.Model(&employee).Where("id = ?", id).Update(emp)
	log.Println("Data Inserted", employee)
	json.NewEncoder(res).Encode(emp)
	defer db.Close()
}

//DeleteEmployee ...
func DeleteEmployee(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	db := db.DbConn()
	employee := models.Employee{}
	deletedEmployee := db.Where("id = ?", id).Delete(employee)
	json.NewEncoder(res).Encode(deletedEmployee)
	defer db.Close()
}
