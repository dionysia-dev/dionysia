package task

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/learn-video/dionysia/internal/model"
)

const (
	defaultVideoCodec      = "avc"
	defaultAudioCodec      = "aac"
	defaultFramerate       = 24
	defaultSegmentDuration = 3
)

type CommandConfig struct {
	DefaultVideoProfiles []model.VideoProfile
	DefaultAudioProfiles []model.AudioProfile
}

func NewDefaultCommandConfig() *CommandConfig {
	return &CommandConfig{
		//nolint:gomnd // values are self explanatory
		DefaultVideoProfiles: []model.VideoProfile{
			{
				Codec:          defaultVideoCodec,
				Bitrate:        500,
				MaxKeyInterval: defaultSegmentDuration * defaultFramerate,
				Framerate:      defaultFramerate,
				Width:          640,
				Height:         360,
			},
			{
				Codec:          defaultVideoCodec,
				Bitrate:        800,
				MaxKeyInterval: defaultSegmentDuration * defaultFramerate,
				Framerate:      defaultFramerate,
				Width:          842,
				Height:         480,
			},
			{
				Codec:          defaultVideoCodec,
				Bitrate:        1200,
				MaxKeyInterval: defaultSegmentDuration * defaultFramerate,
				Framerate:      defaultFramerate,
				Width:          1280,
				Height:         720,
			},
			{
				Codec:          defaultVideoCodec,
				Bitrate:        2500,
				MaxKeyInterval: defaultSegmentDuration * defaultFramerate,
				Framerate:      defaultFramerate,
				Width:          1920,
				Height:         1080,
			},
		},
		//nolint:gomnd // values are self explanatory
		DefaultAudioProfiles: []model.AudioProfile{
			{
				Codec: defaultAudioCodec,
				Rate:  128,
			},
		},
	}
}

type GPACCommandBuilder struct {
	config *CommandConfig
}

func NewGPACCommandBuilder(config *CommandConfig) *GPACCommandBuilder {
	return &GPACCommandBuilder{config: config}
}

func (b *GPACCommandBuilder) Build(id, address, output string, input model.Input) *GPACCommand {
	if len(input.VideoProfiles) == 0 {
		input.VideoProfiles = b.config.DefaultVideoProfiles
	}

	if len(input.AudioProfiles) == 0 {
		input.AudioProfiles = b.config.DefaultAudioProfiles
	}

	return &GPACCommand{
		ID:      id,
		Address: address,
		Output:  output,
		Input:   input,
		Runner: func(program string, args []string) error {
			cmd := exec.Command(program, args...)
			return cmd.Run()
		},
	}
}

type GPACCommand struct {
	ID      string
	Address string
	Output  string
	Input   model.Input
	Runner  func(string, []string) error
}

func (g *GPACCommand) Execute() error {
	args := []string{"-i", fmt.Sprintf("%s/%s", g.Address, g.ID)}

	for i, v := range g.Input.VideoProfiles {
		bitrate := fmt.Sprintf("b=%dk", v.Bitrate)
		profileFlag := "@@"

		if i == 0 {
			profileFlag = "@"
		}

		args = append(args, profileFlag, bitrate)
	}

	for _, a := range g.Input.AudioProfiles {
		args = append(args, "@@", fmt.Sprintf("c=aac:b=%dk", a.Rate))
	}

	args = append(args, "-o", fmt.Sprintf("%s/%s/index.m3u8:segdur=2:dmode=dynamic:profile=live:muxtype=ts", g.Output, g.ID))

	log.Printf("Executing gpac command: %s", strings.Join(args, " "))

	return g.Runner("gpac", args)
}

func NewGPACCommand(id, address, output string, input model.Input) *GPACCommand {
	runner := func(program string, args []string) error {
		cmd := exec.Command(program, args...)
		return cmd.Run()
	}

	return &GPACCommand{
		ID:      id,
		Address: address,
		Output:  output,
		Input:   input,
		Runner:  runner,
	}
}
