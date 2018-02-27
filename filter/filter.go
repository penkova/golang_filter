package filter

import (
	"encoding/json"
	"fmt"
	"github.com/user/golang_filter/db"
	"net/http"
	"strconv"
)

func handleError(err error, message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(message, err)))
}

// GetCarsFilter -- Filtering cars by some parameters.
func GetCarsFilter(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Filter for Cars by model: \n \t"))

	valuesQuery := req.URL.Query()
	fmt.Println("values:", valuesQuery)

	valueModel := valuesQuery.Get("model")
	modelsCar := valuesQuery["model"]

	valueAgeCarString := valuesQuery.Get("age")
	valueAgeIntCar, _ := strconv.Atoi(valueAgeCarString)
	ageCar := valuesQuery["age"]
	valueAgeCar := []int{}

	for _, v := range ageCar {
		valueAge, err := strconv.Atoi(v)
		if err != nil {
			handleError(err, "\t\t Failed to convertation: %v", w)
			return
		}
		valueAgeCar = append(valueAgeCar, valueAge)
	}

	//valuePriceCarString := valuesQuery.Get("price")
	//valuePriceIntCar, _ := strconv.Atoi(valuePriceCarString)

	fmt.Fprintf(w, "Model: %s \n \t", modelsCar)
	fmt.Fprintf(w, "Age Int: %d \n \t", valueAgeCar)
	//fmt.Fprintf(w, "Price Int: %d \n \t", valuePriceIntCar)

	if valueModel != "" && valueAgeIntCar != 0 {
		// Filtering by model and age(once) car
		fmt.Println("Filtering by model and age. Params model and age:", modelsCar, valueAgeCar)
		rs, err := db.FindCarAgeName(valueAgeCar, modelsCar)
		if err != nil {
			handleError(err, "\n\t\t Failed to read database: %v", w)
			return
		}
		json.NewEncoder(w).Encode(rs)
		return
	} else if valueAgeIntCar == 0 {
		// Filtering by model
		fmt.Println("Filtering by model. Params model:", modelsCar)
		rs, err := db.FindCarModel(modelsCar)
		if err != nil {
			handleError(err, "\t\t Failed to read database: %v", w)
			return
		}
		json.NewEncoder(w).Encode(rs)
		return
	} else if valueModel == "" {
		// Filtering by age car
		fmt.Println("Filtering by age. Params age:", modelsCar)

		rs, err := db.FindCarAge(valueAgeCar)
		if err != nil {
			handleError(err, "\n\t\t Failed to read database: %v", w)
			return
		}
		json.NewEncoder(w).Encode(rs)
		return
	}
}

// GetPeopleFilter -- Filtering people by some parameters.
func GetPeopleFilter(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Filter for People by name: \n \t"))

	valuesQuery := req.URL.Query()
	fmt.Println("values:", valuesQuery)

	valueName := valuesQuery.Get("name")
	namePerson := valuesQuery["name"]

	valueAgePersonString := valuesQuery.Get("age")
	valueAgePersonInt, _ := strconv.Atoi(valueAgePersonString)

	agePerson := valuesQuery["age"]
	valueAgePerson := []int{}
	for _, v := range agePerson {
		valueAge, err := strconv.Atoi(v)
		if err != nil {
			handleError(err, "\t\t Failed to convertation: %v", w)
			return
		}
		valueAgePerson = append(valueAgePerson, valueAge)
	}

	fmt.Fprintf(w, "Name: %s \n \t", valueName)
	fmt.Fprintf(w, "Age Int: %d \n \t", valueAgePersonInt)

	if valueName != "" && valueAgePersonInt != 0 {
		// Filtering by model and age(once) person
		fmt.Println("Filtering by model and age. Params model and age:", namePerson, valueAgePerson)
		rs, err := db.FindPeopleAgeName(valueAgePerson, namePerson)
		if err != nil {
			handleError(err, "\n\t\t Failed to read database: %v", w)
			return
		}
		json.NewEncoder(w).Encode(rs)
		return
	} else if valueAgePersonInt == 0 {
		// Filtering by name
		fmt.Println("Filtering by model. Params model:", namePerson)
		rs, err := db.FindPeopleName(namePerson)
		if err != nil {
			handleError(err, "\t\t Failed to read database: %v", w)
			return
		}
		json.NewEncoder(w).Encode(rs)
		return
	} else if valueName == "" {
		// Filtering by age person
		fmt.Println("Filtering by age. Params age:", valueAgePerson)

		rs, err := db.FindPeopleAge(valueAgePerson)
		if err != nil {
			handleError(err, "\n\t\t Failed to read database: %v", w)
			return
		}
		json.NewEncoder(w).Encode(rs)
		return
	}
}

func GetFilter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You can filter by cars or people! \n \t"))
}
