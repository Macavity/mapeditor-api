package models

import "Macavity/mapeditor-server/server/TileMaps"

type Layer struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Type    TileMaps.LayerType `validate:"required" json:"type"`
	X       uint               `gorm:"default:0" json:"x"`
	Y       uint               `gorm:"default:0" json:"y"`
	Z       uint               `gorm:"default:0" json:"z"`
	Data    []uint             `json:"data"`
	Visible bool               `gorm:"default:true" json:"visible"`
	Opacity uint               `gorm:"default:1" json:"opacity"`
	Width   uint               `gorm:"not null" validate:"required" json:"width"`
	Height  uint               `gorm:"not null" validate:"required" json:"height"`
}
