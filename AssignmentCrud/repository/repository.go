package repository

import (
	"employeeeDirectory/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type EmployeeRepo struct {
	employees map[int]models.Employee
}

func NewEmployeeRepo() *EmployeeRepo {
	return &EmployeeRepo{employees: make(map[int]models.Employee)}
}

func (r *EmployeeRepo) CreateEmployee(w http.ResponseWriter, req *http.Request) {

	var emp models.Employee

	if err := json.NewDecoder(req.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}

	fmt.Println(emp)

	if _, exists := r.employees[emp.ID()]; exists {
		fmt.Println("Employee Already Exists")
		http.Error(w, "Invalid Request body", http.StatusUnprocessableEntity)
		return
	}

	r.employees[emp.ID()] = emp

	r.ListAllEmployees()

	w.WriteHeader(http.StatusCreated)

}

func (r *EmployeeRepo) GetEmployee(w http.ResponseWriter, req *http.Request, id int) {
	if val, exists := r.employees[id]; !exists {
		http.Error(w, "No Employee Found", http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(val)
	}

}

func (r *EmployeeRepo) UpdateEmployee(w http.ResponseWriter, req *http.Request, id int) {
	var emp models.Employee

	if err := json.NewDecoder(req.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if _, exists := r.employees[id]; !exists {
		http.Error(w, "No Employee Found to Update", http.StatusNotFound)
		return
	}

	emp.EmployeeID = id
	r.employees[id] = emp
	r.ListAllEmployees()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Employee updated: %+v", emp)
}

func (r *EmployeeRepo) DeleteEmployee(w http.ResponseWriter, req *http.Request, id int) {
	if _, exists := r.employees[id]; !exists {
		http.Error(w, "No Employee Found to Delete", http.StatusNotFound)
		return
	}

	delete(r.employees, id)
	r.ListAllEmployees()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Employee deleted with ID: %d", id)
}

func (r *EmployeeRepo) ListAllEmployees() {
	fmt.Println(r.employees)
}
