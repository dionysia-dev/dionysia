package service_test

import (
	"testing"

	"github.com/dionysia-dev/dionysia/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestGPACCommandExecute(t *testing.T) {
	input := service.Input{
		VideoProfiles: []service.VideoProfile{{Bitrate: 1000}, {Bitrate: 2000}},
		AudioProfiles: []service.AudioProfile{{Bitrate: 128}},
	}
	cmd := service.NewGPACCommand("1234", "rtmp://localhost", "/output", input)

	mockRunner := func(program string, args []string) error {
		assert.Equal(t, "gpac", program)
		expectedArgs := []string{
			"-i", "rtmp://localhost/1234",
			"@", "c=avc:b=1000k",
			"@@", "c=avc:b=2000k",
			"@@", "c=aac:b=128k",
			"@", "@1", "@2",
			"-o", "/output/1234/index.m3u8:segdur=3:dmode=dynamic:profile=live:muxtype=ts",
		}
		assert.Equal(t, expectedArgs, args)
		return nil
	}

	cmd.Runner = mockRunner
	err := cmd.Execute()
	assert.NoError(t, err)
}
