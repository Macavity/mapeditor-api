package database

import (
	"Macavity/mapeditor-server/server/Users/models"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Client *gorm.DB
var err error

func Connect(User, Password, Port, Host, Name string) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", Host, Port, User, Name, Password)
	Client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Warnf("Cannot connect to database")
		log.Fatal("This is the error:", err)
		return
	}

	log.Infof("We are connected to the database")
}

func MigrateDatabase() {
	err = Client.AutoMigrate(&models.User{})
	if err != nil {
		log.Println("Database: Migration completed.")
	} else {
		log.Fatalln("Database: Migration failed.", err)
	}
}
