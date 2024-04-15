package task_test

import (
	"testing"

	"github.com/learn-video/dionysia/internal/model"
	"github.com/learn-video/dionysia/internal/task"
	"github.com/stretchr/testify/assert"
)

func TestGPACCommandExecute(t *testing.T) {
	input := model.Input{
		VideoProfiles: []model.VideoProfile{{Bitrate: 1000}, {Bitrate: 2000}},
		AudioProfiles: []model.AudioProfile{{Rate: 128}},
	}
	cmd := task.NewGPACCommand("1234", "rtmp://localhost", "/output", input)

	mockRunner := func(program string, args []string) error {
		assert.Equal(t, "gpac", program)
		expectedArgs := []string{
			"-i", "rtmp://localhost/1234",
			"@", "b=1000k",
			"@@", "b=2000k",
			"@@", "c=aac:b=128k",
			"-o", "/output/1234/index.m3u8:segdur=2:dmode=dynamic:profile=live:muxtype=ts",
		}
		assert.Equal(t, expectedArgs, args)
		return nil
	}

	cmd.Runner = mockRunner
	err := cmd.Execute()
	assert.NoError(t, err)
}
