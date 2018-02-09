package db

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type People struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
	Age  int           `json:"age" bson:"age"`
	//car []Cars
	// id = car_id
}

type Cars struct {
	//car_id bson.ObjectId 	`json:"id" bson:"_id"`
	Id    bson.ObjectId `json:"id" bson:"_id"`
	Model string        `json:"model" bson:"model"`
	Age   int           `json:"age" bson:"age"`
	Price int           `json:"price" bson:"price"`
}

//type Filter struct {
//	Id     bson.ObjectId `json:"id" bson:"_id,omitempty"`
//	Offset string        `json:"offset"`
//	Limit  string        `json:"limit"`
//	//Skip int			`json:"skip"`
//	Name  string `json:"name"`
//	Model string `json:"model"`
//	Age   int    `json:"age"`
//	Price int    `json:"price"`
//}

var db *mgo.Database

func init() {
	session, err := mgo.Dial("localhost:27018")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db = session.DB("crudmongo")
}

// -- Connection to cars and people collections
func CollectionPerson() *mgo.Collection {
	return db.C("peopledb")
}
func CollectionCars() *mgo.Collection {
	return db.C("carsdb")
}

// -- Finding Cars by some parameters
func FindCarModel(model ...string) ([]Cars, error) {
	res := []Cars{}
	if err := CollectionCars().Find(bson.M{"model": bson.M{"$in": model}}).All(&res); err != nil {
		return nil, err
	}
	return res, nil
}
func FindCarAge(carAge ...int) ([]Cars, error) {
	res := []Cars{}
	if err := CollectionCars().Find(bson.M{"age": bson.M{"$in": carAge}}).All(&res); err != nil {
		return nil, err
	}
	return res, nil
}
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

//	-- Finding People by some parameters
func FindPeopleName(name ...string) ([]People, error) {
	res := []People{}
	if err := CollectionCars().Find(bson.M{"name": bson.M{"$in": name}}).All(&res); err != nil {
		return nil, err
	}
	return res, nil
}
func FindPeopleAge(personAge ...int) ([]People, error) {
	res := []People{}
	if err := CollectionCars().Find(bson.M{"age": bson.M{"$in": personAge}}).All(&res); err != nil {
		return nil, err
	}
	return res, nil
}
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

// -- GetAll returns all items from the database.
func GetAllPerson() ([]People, error) {
	res := []People{}

	if err := CollectionPerson().Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}
func GetAllCars() ([]Cars, error) {
	res := []Cars{}

	if err := CollectionCars().Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

// -- GetOne returns a single item from the database.
func GetOneCar(id string) (*Cars, error) {
	res := Cars{}

	if err := CollectionCars().FindId(bson.ObjectIdHex(id)).One(&res); err != nil {
		return nil, err
	}
	return &res, nil
}
func GetOnePerson(id string) (*People, error) {
	res := People{}
	fmt.Println(id)

	if err := CollectionPerson().FindId(bson.ObjectIdHex(id)).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

// -- Create    inserts an item to the database.
func CreateOnePerson(person People) error {
	return CollectionPerson().Insert(person)
}
func CreateOneCar(car Cars) error {
	return CollectionCars().Insert(car)
}

// -- Remove deletes an item from the database
func RemovePerson(id string) error {
	return CollectionPerson().Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
func RemoveCar(id string) error {
	return CollectionCars().Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}

// -- Update
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
