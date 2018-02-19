package db

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
// -- People represents database entity.
type People struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
	Age  int           `json:"age" bson:"age"`
	//car []Cars
	// id = car_id
}
// -- Cars represents database entity.
type Cars struct {
	//car_id bson.ObjectId 	`json:"id" bson:"_id"`
	Id    bson.ObjectId `json:"id" bson:"_id"`
	Model string        `json:"model" bson:"model"`
	Age   int           `json:"age" bson:"age"`
	Price int           `json:"price" bson:"price"`
}

var db *mgo.Database

// -- Establish a connection to MongoDB database.
func init() {
	//Getting a session
	session, err := mgo.Dial("mongodb:27017")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db = session.DB("crudmongo")
}

// -- Connection to people collections
func CollectionPerson() *mgo.Collection {
	return db.C("peopledb")
}
// -- Connection to cars collections
func CollectionCars() *mgo.Collection {
	return db.C("carsdb")
}

// -- Finding person by name parameter..
func FindPeopleName(name []string) ([]People, error) {
	res := []People{}
	if err := CollectionCars().Find(bson.M{"name": bson.M{"$in": name}}).All(&res); err != nil {
		return nil, err
	}
	return res, nil
}
// -- Finding People by age parameter.
func FindPeopleAge(personAge []int) ([]People, error) {
	res := []People{}
	if err := CollectionCars().Find(bson.M{"age": bson.M{"$in": personAge}}).All(&res); err != nil {
		return nil, err
	}
	return res, nil
}
// -- Finding People by name and age parameters.
func FindPeopleAgeName(personAge []int, name []string) ([]People, error) {
	res := []People{}
	if err := CollectionCars().Find(bson.M{
		"model": bson.M{"$in": name},
		"age":   bson.M{"$in": personAge},
	}).All(&res); err != nil {
		return nil, err
	}
	return res, nil
}

// -- Finding Cars by model parameter.
func FindCarModel(model []string) ([]Cars, error) {
	res := []Cars{}
	if err := CollectionCars().Find(bson.M{"model": bson.M{"$in": model}}).All(&res); err != nil {
		return nil, err
	}
	return res, nil
}
// -- Finding Cars by age parameter.
func FindCarAge(carAge []int) ([]Cars, error) {
	res := []Cars{}
	if err := CollectionCars().Find(bson.M{"age": bson.M{"$in": carAge}}).All(&res); err != nil {
		return nil, err
	}
	return res, nil
}
// -- Finding Cars by model and age parameters.
func FindCarAgeName(carAge []int, model []string) ([]Cars, error) {
	res := []Cars{}
	if err := CollectionCars().Find(bson.M{
		"model": bson.M{"$in": model},
		"age":   bson.M{"$in": carAge},
	}).All(&res); err != nil {
		return nil, err
	}
	return res, nil
}

// -- Returns all person from the database.
func GetAllPerson() ([]People, error) {
	res := []People{}

	if err := CollectionPerson().Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}
// -- Returns all cars from the database.
func GetAllCars() ([]Cars, error) {
	res := []Cars{}

	if err := CollectionCars().Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

// -- Returns a single person from the database.
func GetOnePerson(id string) (*People, error) {
	res := People{}
	fmt.Println(id)

	if err := CollectionPerson().FindId(bson.ObjectIdHex(id)).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}
// -- Returns a single car from the database.
func GetOneCar(id string) (*Cars, error) {
	res := Cars{}

	if err := CollectionCars().FindId(bson.ObjectIdHex(id)).One(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// -- Create  person to the database.
func CreateOnePerson(person People) error {
	return CollectionPerson().Insert(person)
}
// -- Create car to the database.
func CreateOneCar(car Cars) error {
	return CollectionCars().Insert(car)
}

// -- Remove person from the database
func RemovePerson(id string) error {
	return CollectionPerson().Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
// -- Remove car from the database
func RemoveCar(id string) error {
	return CollectionCars().Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}

// -- Update person from the database by id.
func UpdateOnePerson(id string, p *People) (*People, error) {
	res := People{}
	if err := CollectionPerson().Update(bson.M{
		"_id": bson.ObjectIdHex(id),
	}, bson.M{
		"$set": bson.M{
			"name": p.Name,
			"age":  p.Age,
		}}); err != nil {
		return nil, err
	}
	return &res, nil
}
// -- Update car from the database y id.
func UpdateOneCar(id string, c *Cars) (*Cars, error) {
	if err := CollectionCars().Update(bson.M{
		"_id": bson.ObjectIdHex(id),
	}, bson.M{
		"$set": bson.M{
			"model": c.Model,
			"age":   c.Age,
			"price": c.Price,
		}}); err != nil {
		return nil, err
	}
	return c, nil

}
