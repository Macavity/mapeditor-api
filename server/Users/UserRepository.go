package Users

import (
	"Macavity/mapeditor-server/server/Users/models"
	"Macavity/mapeditor-server/server/database"
	"fmt"
)

func FindAllUsers() (*[]models.User, error) {
	var err error
	var users []models.User
	err = database.Client.Debug().Limit(100).Find(&users).Error
	if err != nil {
		return &[]models.User{}, err
	}
	return &users, err
}

func CreateUser(user *models.User) {
	database.Client.Debug().Create(&user)
}

func FindUserById(userId uint64) (user models.User, err error) {
	database.Client.First(&user, userId)

	if user.ID == 0 {
		err = fmt.Errorf("user with ID %d not found", userId)
	}

	return user, err
}
