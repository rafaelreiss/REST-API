package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

// User field definition.
type User struct {
	Id        string `json:"id" bson:"_id,omitempty"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Email     string `json:"email" bson:"email"`
}

var db *mgo.Database

// Connects to the database.
func init() {
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		log.Fatalf("Failed connecting into the db: %v", err)
	}

	db = session.DB("godb_test")
}

// Creates <Users> document.
func collection() *mgo.Collection {
	return db.C("Users")
}

// Return all the users in the Users document.
func GetAllUsers() ([]User, error) {
	user := []User{}

	if err := collection().Find(bson.M{}).All(&user); err != nil {
		return nil, err
	}
	return user, nil
}

// Creates an User
func Create(user User) error {
	return collection().Insert(user)
}

// Search for an user by his <id>.
// It returns an User object
func FindById(id string) (*User, error) {
	user := User{}

	if err := collection().Find(bson.M{"_id": id}).One(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// Deletes an User from the document
func Delete(id string) error {
	return collection().Remove(bson.M{"_id": id})
}
