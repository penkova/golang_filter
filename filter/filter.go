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
	//if valueAgeCar != "" {
	//	carAge := valuesQuery["age"]
	//	fmt.Println("GET params age:", carAge)
	//
	//	for _ , value := range carAge {
	//		valueAgeIntCar, _:= strconv.Atoi(value)
	//		rs, err := db.FindCarAge(valueAgeIntCar)
	//		if err != nil {
	//			handleError(err, "\n\t\t Failed to read database: %v", w)
	//			return
	//		}
	//	json.NewEncoder(w).Encode(rs)
	//	}
	//	return
	//}
	if valueAgeIntCar != 0 {
		ageCar := valuesQuery["age"]
		for i := 0; i < len(ageCar); i++ {
			fmt.Printf("\n value %s", ageCar[i])
			valueAge, _ := strconv.Atoi(ageCar[i])
			fmt.Println("GET params carAge:", ageCar)
			rs, err := db.FindCarAge(valueAge)
			if err != nil {
				handleError(err, "\n\t\t Failed to read database: %v", w)
				return
			}
			json.NewEncoder(w).Encode(rs)
			return
			//ageCar := valuesQuery["age"]
			//fmt.Println("GET params carAge:", ageCar)
			//rs, err := db.FindCarAge(ageCar...)
			//if err != nil {
			//	handleError(err, "\n\t\t Failed to read database: %v", w)
			//	return
			//}
			//json.NewEncoder(w).Encode(rs)
			//return
		}
	}

	//notice := "Query string is:" + req.Method + req.URL.String()
	//log.Println(notice)
	//valuesFromQuery := req.URL.Query()
	//param1 := valuesFromQuery.Get("name")
	//if param1 != "" {
	//	// ... process it, will be the first (only) if multiple were given
	//	// note: if they pass in like ?param1=&param2= param1 will also be "" :|
	//}
	//param1s := valuesFromQuery["param1"]
	//if param1s != nil {
	//	// ... process them ... or you could just iterate over them without a nil check
	//	// [if none were present, will be a nil array, which golang treats as an empty array]
	//	// this way you can also tell if they passed in the parameter as the empty string
	//	// it will be an element of the array that is the empty string
	//}
}

func GetPeopleFilter(w http.ResponseWriter, req *http.Request) {

}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET params were:", r.URL.Query())

	// if only one expected
	param1 := r.URL.Query().Get("param1")
	if param1 != "" {
		// ... process it, will be the first (only) if multiple were given
		// note: if they pass in like ?param1=&param2= param1 will also be "" :|
	}

	// if multiples possible, or to process empty values like param1 in
	// ?param1=&param2=something
	param1s := r.URL.Query()["param1"]
	if param1s != nil {
		// ... process them ... or you could just iterate over them without a nil check
		// [if none were present, will be a nil array, which golang treats as an empty array]
		// this way you can also tell if they passed in the parameter as the empty string
		// it will be an element of the array that is the empty string
	}
}
