package models

import (
	"Macavity/mapeditor-server/server/TileMaps/models"
	"gorm.io/gorm"
	"time"
)

// use a single instance of Validate, it caches struct info
//var validate *validator.Validate

type User struct {
	ID        uint64         `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Name      string         `json:"name" gorm:"unique" validate:"required"`
	Email     string         `json:"email" gorm:"unique"`

	TileMaps []models.TileMap `json:"tileMaps"`
}

type CreateUserDTO struct {
	Name  string
	Email string
}
