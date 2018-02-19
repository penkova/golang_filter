package main

import (
	"github.com/gorilla/mux"
	"github.com/user/mongo_filter/api"
	"github.com/user/mongo_filter/filter"
	"net/http"
)

func main() {
	// Create routes
	r := mux.NewRouter()

	// --Creates a controller for cars endpoint
	// Request for people
	r.HandleFunc("/people", api.GetPeople).Methods("GET")
	r.HandleFunc("/people/{id}", api.GetPerson).Methods("GET")
	r.HandleFunc("/people", api.CreatePerson).Methods("POST")
	r.HandleFunc("/people/{id}", api.UpdatePerson).Methods("PUT")
	r.HandleFunc("/people/{id}", api.DeletePerson).Methods("DELETE")

	// Request for cars
	r.HandleFunc("/cars", api.GetAllCars).Methods("GET")
	r.HandleFunc("/cars/{id}", api.GetCar).Methods("GET")
	r.HandleFunc("/cars", api.CreateCar).Methods("POST")
	r.HandleFunc("/cars/{id}", api.UpdateCar).Methods("PUT")
	r.HandleFunc("/cars/{id}", api.DeleteCar).Methods("DELETE")

	// Request for filtering items
	r.HandleFunc("/filter/people", filter.GetPeopleFilter).Methods("GET")
	r.HandleFunc("/filter/cars", filter.GetCarsFilter).Methods("GET")
	r.HandleFunc("/filter", filter.GetFilter).Methods("GET")

	http.ListenAndServe(":3030", r)
}
