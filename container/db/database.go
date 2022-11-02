package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// no need to add a package just to add db inside
// you can directly have it inside container

// instead of adding a var here to save the db instance
// you can return one in the below method to be used in the project

var Instance *gorm.DB
var dbError error

// Connect in case we added this directly under container we need a better name for it
func Connect(connectioString string) {
	//why passing the connection string to the method not adding it here
	//it will be constant for the whole project
	Instance, dbError = gorm.Open(postgres.Open(connectioString), &gorm.Config{})
	if dbError != nil {
		log.Fatal("error configuring the database: ", dbError)
	}
	log.Println("Hey! You successfully connected to your db.")
}
