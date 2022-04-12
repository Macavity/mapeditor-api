package TileMaps

type Dimension struct {
	Width  uint `gorm:"not null" validate:"required" json:"width"`
	Height uint `gorm:"not null" validate:"required" json:"height"`
}
