package model

import "github.com/google/uuid"

type Input struct {
	ID     uuid.UUID      `json:"id"`
	Name   string         `json:"name" validate:"required"`
	Format string         `json:"format" validate:"required"`
	Audio  []AudioProfile `json:"audio"`
	Video  []VideoProfile `json:"video"`
}

type AudioProfile struct {
	Rate  int    `json:"rate"`
	Codec string `json:"codec"`
}

type VideoProfile struct {
	Codec          string `json:"codec"`
	Bitrate        int    `json:"bitrate"`
	MaxKeyInterval int    `json:"max_key_interval"`
	Framerate      int    `json:"framerate"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
}
