package service

import (
	"net/http"
)

type EmployeeService interface {
	CreateEmployee(w http.ResponseWriter, r *http.Request) //Create
	GetEmployee(w http.ResponseWriter, r *http.Request, id int)
	UpdateEmployee(w http.ResponseWriter, r *http.Request, id int)
	DeleteEmployee(w http.ResponseWriter, r *http.Request, id int)
	ListAllEmployees()
}
