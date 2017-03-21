package core

import (
	"fmt"
	"log"
	"os"

	. "github.com/jonhmchan/boilerplate/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Storage StorageStruct

type StorageStruct struct {
	DB *gorm.DB
}

func (s *StorageStruct) Start() {
	initDatabase()
}

func initDatabase() {
	log.Println("START: DATABASE INITIALIZING")
	var err error
	Storage.DB, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}

	Storage.DB.LogMode(Config.Local())

	Storage.DB.AutoMigrate(
		&User{},
	)

	log.Println("END: DATABASE INITIALIZING")
}
