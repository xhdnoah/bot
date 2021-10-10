package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *Database

type Database struct {
	Self *gorm.DB
}

func (db *Database) Init() {
	_db, err := gorm.Open(sqlite.Open("joke.db"), &gorm.Config{})
	DB = &Database{Self: _db}
	if err != nil {
		log.Printf("Database connection failed")
	}
	_db.AutoMigrate(&JokeModel{})
}
