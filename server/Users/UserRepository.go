package Users

import (
	"Macavity/mapeditor-server/server/Users/models"
	"Macavity/mapeditor-server/server/database"
	"Macavity/mapeditor-server/server/logwrapper"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
)

func FindAllUsers() (*[]models.User, error) {
	var err error
	var users []models.User
	logger := logwrapper.NewDebugLogger()
	logger.Debugln("FindAllUsers")
	client := database.NewConnection()

	err = client.Model(&models.User{}).Limit(20).Find(&users).Error

	if err != nil {
		logger.Errorln(err)
		return &[]models.User{}, err
	}

	return &users, err
}

func MigrateUserTable() {
	client := database.NewConnection()
	logger := logwrapper.NewLogger(logrus.ErrorLevel)
	logger.Debugln("[Users.Migrate] Start")

	if client.Migrator().HasTable(&models.User{}) == false {
		logger.Debugln("[Users.Migrate] Create Table")
		err := client.Migrator().CreateTable(&models.User{})
		if err != nil {
			logger.Errorln("[Users.Migrate] Table creation failed", err)
			return
		}
	}

	err := client.AutoMigrate(&models.User{})
	if err != nil {
		logger.Errorln("[Users.Migrate] AutoMigrate failed", err)
		return
	}
}

func CreateUser(user *models.User) (err error) {
	client := database.NewConnection()
	if err != nil {
		log.Println(err)
		return err
	}

	err = client.Table("users").Create(user).Error

	return err
}

func FindUserById(userId uint64) (user models.User, err error) {
	database.Client.First(&user, userId)

	if user.ID == 0 {
		err = fmt.Errorf("user with ID %d not found", userId)
	}

	return user, err
}
