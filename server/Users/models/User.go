package models

import (
	"Macavity/mapeditor-server/server/TileMaps/models"
	"gorm.io/gorm"
)

// use a single instance of Validate, it caches struct info
//var validate *validator.Validate

type User struct {
	gorm.Model

	Name  string `json:"name" gorm:"unique" validate:"required"`
	Email string `json:"email" gorm:"unique"`

	TileMaps []models.TileMap
}
