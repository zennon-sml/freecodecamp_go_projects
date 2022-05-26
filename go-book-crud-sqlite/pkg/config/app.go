package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB// set db to a variable of typei gorm
)

func Connect() {
	d, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})// opens a connection with the database 
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db // return the database
}
