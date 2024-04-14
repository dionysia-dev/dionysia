package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Input struct {
	gorm.Model
	ID            uuid.UUID      `gorm:"primary_key;type:uuid"`
	Name          string         `json:"name" validate:"required"`
	Format        string         `json:"format" validate:"required"`
	AudioProfiles []AudioProfile `json:"audio" gorm:"foreignKey:InputID"`
	VideoProfiles []VideoProfile `json:"video" gorm:"foreignKey:InputID"`
}

type AudioProfile struct {
	gorm.Model
	InputID uuid.UUID
	Rate    int    `json:"rate"`
	Codec   string `json:"codec"`
}

type VideoProfile struct {
	gorm.Model
	InputID        uuid.UUID
	Codec          string `json:"codec"`
	Bitrate        int    `json:"bitrate"`
	MaxKeyInterval int    `json:"max_key_interval"`
	Framerate      int    `json:"framerate"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
}
