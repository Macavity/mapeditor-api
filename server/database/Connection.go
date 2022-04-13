package database

import (
	"Macavity/mapeditor-server/server/logwrapper"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var Client *gorm.DB
var err error

func NewConnection() *gorm.DB {
	client := Connect(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	return client
}

func Connect(User, Password, Port, Host, Name string) *gorm.DB {
	logger := logwrapper.NewDebugLogger()
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", Host, Port, User, Name, Password)

	Client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Warnln("[Database/Connection] Cannot connect to database")
		logger.Fatalln(err)
		return nil
	}

	return Client
}

//func MigrateDatabase() {
//	client := NewConnection()
//	logger := logwrapper.NewDebugLogger()
//	//err = client.AutoMigrate(&userModels.User{}, &tileMapModels.TileMap{}, tileMapModels.Layer{})
//	if err != nil {
//		logger.Println("[Database/Connection] Migration completed.")
//	} else {
//		logger.Fatalln("[Database/Connection] Migration failed.", err)
//	}
//}
