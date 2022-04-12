package database

import (
	"Macavity/mapeditor-server/server/Users"
	"Macavity/mapeditor-server/server/Users/models"
	"Macavity/mapeditor-server/server/logwrapper"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var Client *gorm.DB
var err error

func NewConnection() *gorm.DB {
	client := Connect(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	Users.MigrateUserTable()

	return client
}

func Connect(User, Password, Port, Host, Name string) *gorm.DB {
	logger := logwrapper.NewDebugLogger()
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", Host, Port, User, Name, Password)
	log.Println("Database.Connect")
	Client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Warnln("Cannot connect to database")
		logger.Fatalln(err)
		return nil
	}

	logger.Debugln("We are connected to the database")

	return Client
}

func MigrateDatabase(Logger *logwrapper.StandardLogger) {
	err = Client.AutoMigrate(&models.User{})
	if err != nil {
		Logger.Println("Database: Migration completed.")
	} else {
		Logger.Fatalln("Database: Migration failed.", err)
	}
}
