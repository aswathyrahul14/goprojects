package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Employee struct
type Employee struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var employees []Employee

// Get all employees
func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

// Get single employee record
func getEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range employees {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Employee{})
}

// Add new employee
func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)
	employee.ID = strconv.Itoa(rand.Intn(100000000))
	employees = append(employees, employee)
	json.NewEncoder(w).Encode(employee)
}

// Update employee of given id
func updateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range employees {
		if item.ID == params["id"] {
			employees = append(employees[:index], employees[index+1:]...)
			var employee Employee
			_ = json.NewDecoder(r.Body).Decode(&employee)
			employee.ID = params["id"]
			employees = append(employees, employee)
			json.NewEncoder(w).Encode(employee)
			return
		}
	}
}

// Delete employee record of given id
func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range employees {
		if item.ID == params["id"] {
			employees = append(employees[:index], employees[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(employees)
}

func main() {

	r := mux.NewRouter()
	employees = append(employees, Employee{ID: "1", Firstname: "Aswathy", Lastname: "BS"})
	employees = append(employees, Employee{ID: "2", Firstname: "Rahul", Lastname: "AV"})
	r.HandleFunc("/employees", getEmployees).Methods("GET")
	r.HandleFunc("/employees/{id}", getEmployee).Methods("GET")
	r.HandleFunc("/employees", createEmployee).Methods("POST")
	r.HandleFunc("/employees/{id}", updateEmployee).Methods("PUT")
	r.HandleFunc("/employees/{id}", deleteEmployee).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
