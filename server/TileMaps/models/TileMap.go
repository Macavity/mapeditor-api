package models

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type TileMap struct {
	ID         uint64         `gorm:"primaryKey" json:"-"`
	UUID       uuid.UUID      `gorm:"index" json:"uuid"`
	UserID     uint64         `json:"userId"`
	Width      uint           `gorm:"not null" validate:"required" json:"width"`
	Height     uint           `gorm:"not null" validate:"required" json:"height"`
	Name       string         `gorm:"size:255;not null" validate:"required" json:"name"`
	TileWidth  uint           `validate:"required" json:"tileWidth"`
	TileHeight uint           `validate:"required" json:"tileHeight"`
	Layers     []Layer        `json:"layers"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type CreateTileMapDTO struct {
	UserID uint64 `json:"userId"`
	Width  uint   `gorm:"not null" validate:"required" json:"width"`
	Height uint   `gorm:"not null" validate:"required" json:"height"`
	Name   string `gorm:"size:255;not null" validate:"required" json:"name"`
}

func (m *TileMap) Validate() {
	v := validator.New()
	a := TileMap{}
	err := v.Struct(a)

	for _, e := range err.(validator.ValidationErrors) {
		fmt.Println(e)
	}
}
