package api

import "github.com/google/uuid"

type InputData struct {
	ID            uuid.UUID          `json:"id" swaggerignore:"true"`
	Name          string             `json:"name" validate:"required"`
	Format        string             `json:"format" validate:"required"`
	AudioProfiles []AudioProfileData `json:"audio_profiles"`
	VideoProfiles []VideoProfileData `json:"video_profiles"`
}

type AudioProfileData struct {
	Codec   string `json:"codec"`
	Bitrate int    `json:"bitrate"`
}

type VideoProfileData struct {
	Codec          string `json:"codec"`
	Bitrate        int    `json:"bitrate"`
	MaxKeyInterval int    `json:"max_key_interval"`
	Framerate      int    `json:"framerate"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
}

type IngestAuthData struct {
	Path   uuid.UUID `json:"path" validate:"required"`
	Action string    `json:"action" validate:"required"`
}
