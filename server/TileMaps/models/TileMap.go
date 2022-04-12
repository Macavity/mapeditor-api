package models

import (
	"Macavity/mapeditor-server/server/TileMaps"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"time"
)

type TileMap struct {
	gorm.Model

	TileMaps.Dimension
	UserID     uint
	Name       string `gorm:"size:255;not null" validate:"required" json:"name"`
	TileWidth  uint   `validate:"required" json:"tileWidth"`
	TileHeight uint   `validate:"required" json:"tileHeight"`
	Layers     []Layer
}

func (m *TileMap) Prepare() {
	m.ID = 0
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
}

func (m *TileMap) Validate() {
	v := validator.New()
	a := TileMap{}
	err := v.Struct(a)

	for _, e := range err.(validator.ValidationErrors) {
		fmt.Println(e)
	}
}
