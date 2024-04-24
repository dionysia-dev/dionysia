package model

import (
	"github.com/google/uuid"
)

type Input struct {
	ID            uuid.UUID `gorm:"primary_key;type:uuid"`
	Name          string
	Format        string
	AudioProfiles []AudioProfile `gorm:"foreignKey:InputID"`
	VideoProfiles []VideoProfile `gorm:"foreignKey:InputID"`
}

type AudioProfile struct {
	ID      uint `gorm:"primaryKey"`
	InputID uuid.UUID
	Input   Input
	Codec   string
	Bitrate int
}

type VideoProfile struct {
	ID             uint `gorm:"primaryKey"`
	InputID        uuid.UUID
	Input          Input
	Codec          string
	Bitrate        int
	MaxKeyInterval int
	Framerate      int
	Width          int
	Height         int
}
