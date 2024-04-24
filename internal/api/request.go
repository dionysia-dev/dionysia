package api

import (
	"github.com/dionysia-dev/dionysia/internal/service"
	"github.com/google/uuid"
)

type InputData struct {
	ID            uuid.UUID          `json:"id" swaggerignore:"true"`
	Name          string             `json:"name" validate:"required"`
	Format        string             `json:"format" validate:"required"`
	AudioProfiles []AudioProfileData `json:"audio_profiles"`
	VideoProfiles []VideoProfileData `json:"video_profiles"`
}

func (i *InputData) ToInput() *service.Input {
	audioProfiles := make([]service.AudioProfile, len(i.AudioProfiles))
	for j, audioProfile := range i.AudioProfiles {
		audioProfiles[j] = service.AudioProfile{
			Codec:   audioProfile.Codec,
			Bitrate: audioProfile.Bitrate,
		}
	}

	videoProfiles := make([]service.VideoProfile, len(i.VideoProfiles))
	for j, videoProfile := range i.VideoProfiles {
		videoProfiles[j] = service.VideoProfile{
			Codec:          videoProfile.Codec,
			Bitrate:        videoProfile.Bitrate,
			MaxKeyInterval: videoProfile.MaxKeyInterval,
			Framerate:      videoProfile.Framerate,
			Width:          videoProfile.Width,
			Height:         videoProfile.Height,
		}
	}

	return &service.Input{
		ID:            i.ID,
		Name:          i.Name,
		Format:        i.Format,
		AudioProfiles: audioProfiles,
		VideoProfiles: videoProfiles,
	}
}

func FromInput(input service.Input) InputData {
	audioProfiles := make([]AudioProfileData, len(input.AudioProfiles))
	for i, audioProfile := range input.AudioProfiles {
		audioProfiles[i] = AudioProfileData{
			Codec:   audioProfile.Codec,
			Bitrate: audioProfile.Bitrate,
		}
	}

	videoProfiles := make([]VideoProfileData, len(input.VideoProfiles))
	for i, videoProfile := range input.VideoProfiles {
		videoProfiles[i] = VideoProfileData{
			Codec:          videoProfile.Codec,
			Bitrate:        videoProfile.Bitrate,
			MaxKeyInterval: videoProfile.MaxKeyInterval,
			Framerate:      videoProfile.Framerate,
			Width:          videoProfile.Width,
			Height:         videoProfile.Height,
		}
	}

	return InputData{
		ID:            input.ID,
		Name:          input.Name,
		Format:        input.Format,
		AudioProfiles: audioProfiles,
		VideoProfiles: videoProfiles,
	}
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

type OriginData struct {
	ID      uuid.UUID `json:"id" validate:"required"`
	Address string    `json:"address" validate:"required"`
}

func (o *OriginData) ToOrigin() service.Origin {
	return service.Origin{
		ID:      o.ID,
		Address: o.Address,
	}
}
