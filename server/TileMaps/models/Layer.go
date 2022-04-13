package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Layer struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	TileMapID uint64    `json:"tileMapId"`
	Name      string    `json:"name"`
	Type      LayerType `validate:"required" json:"type"`
	X         uint      `gorm:"default:0" json:"x"`
	Y         uint      `gorm:"default:0" json:"y"`
	Z         uint      `gorm:"default:0" json:"z"`
	Data      JSON      `json:"data"`
	Visible   bool      `gorm:"default:true" json:"visible"`
	Opacity   uint      `gorm:"default:1" json:"opacity"`
	Width     uint      `gorm:"not null" validate:"required" json:"width"`
	Height    uint      `gorm:"not null" validate:"required" json:"height"`
}

type JSON json.RawMessage

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}
