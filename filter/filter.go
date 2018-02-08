package filter

import (
	"encoding/json"
	"fmt"
	"github.com/user/mongo_filter/db"
	"net/http"
	"strconv"
)

func handleError(err error, message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(message, err)))
}

// Filtering by Car
func GetCarsFilter(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Filter for Cars by model: \n \t"))

	valuesQuery := req.URL.Query()
	fmt.Println("values:", valuesQuery)

	// Filtering by model
	valueModel := valuesQuery.Get("model")

	valueAgeCarString := valuesQuery.Get("age")
	valueAgeIntCar, _ := strconv.Atoi(valueAgeCarString)

	valuePriceCarString := valuesQuery.Get("price")
	valuePriceIntCar, _ := strconv.Atoi(valuePriceCarString)

	fmt.Fprintf(w, "Model: %s \n \t", valueModel)
	fmt.Fprintf(w, "Age Int: %d \n \t", valueAgeIntCar)
	fmt.Fprintf(w, "Price Int: %d \n \t", valuePriceIntCar)

	if valueModel != "" {
		modelsCar := valuesQuery["model"]
		fmt.Println("GET params model:", modelsCar)
		rs, err := db.FindCarModel(modelsCar...)
		if err != nil {
			handleError(err, "\t\t Failed to read database: %v", w)
			return
		}

		json.NewEncoder(w).Encode(rs)
		return
	}
	if  valueModel != "" && valueAgeIntCar != 0 {
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
		rs, err := db.FindCarAge(valueAgeCar...)
		if err != nil {
			handleError(err, "\n\t\t Failed to read database: %v", w)
			return
		}
		json.NewEncoder(w).Encode(rs)
		return
	}


	//if valueAgeIntCar != 0 {
	//	ageCar := valuesQuery["age"]
	//	valueAgeCar := []int{}
	//	for _, v := range ageCar {
	//		valueAge, err := strconv.Atoi(v)
	//		if err != nil {
	//			handleError(err, "\t\t Failed to convertation: %v", w)
	//			return
	//		}
	//		valueAgeCar = append(valueAgeCar, valueAge)
	//	}
	//	rs, err := db.FindCarAge(valueAgeCar...)
	//	if err != nil {
	//		handleError(err, "\n\t\t Failed to read database: %v", w)
	//		return
	//	}
	//	json.NewEncoder(w).Encode(rs)
	//	return
	//}
}

func GetPeopleFilter(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Filter for People by name: \n \t"))

	valuesQuery := req.URL.Query()
	fmt.Println("values:", valuesQuery)

	// Filtering by name
	valueName := valuesQuery.Get("name")

	valueAgePersonString := valuesQuery.Get("age")
	valueAgepersonInt, _ := strconv.Atoi(valueAgePersonString)

	fmt.Fprintf(w, "name: %s \n \t", valueName)
	fmt.Fprintf(w, "Age Int: %d \n \t", valueAgepersonInt)



	if valueName != "" {
		namePerson := valuesQuery["name"]
		fmt.Println("GET params model:", namePerson)
		rs, err := db.FindPeopleName(namePerson...)
		if err != nil {
			handleError(err, "\t\t Failed to read database: %v", w)
			return
		}

		json.NewEncoder(w).Encode(rs)
		return
	}
	if  valueName != "" && valueAgepersonInt != 0 {
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
		rs, err := db.FindPeopleAge(valueAgePerson...)
		if err != nil {
			handleError(err, "\n\t\t Failed to read database: %v", w)
			return
		}
		json.NewEncoder(w).Encode(rs)
		return
	}


}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET params were:", r.URL.Query())

	param1 := r.URL.Query().Get("param1")
	if param1 != "" {
	}

	param1s := r.URL.Query()["param1"]
	if param1s != nil {

	}
}
