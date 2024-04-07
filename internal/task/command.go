package task

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type CommandExecutor interface {
	Execute() error
}

type GPACCommand struct {
	ID      string
	Address string
	Output  string
}

func (g *GPACCommand) Execute() error {
	args := []string{
		"-i", fmt.Sprintf("%s/%s", g.Address, g.ID),
		"@",
		"c=avc:b=1m",
		"@@", "c=avc:b=2m",
		"@@", "c=aac:b=44k",
		"@",
		"@1", "@2",
		"-o", fmt.Sprintf("%s/%s/index.m3u8:segdur=2:dmode=dynamic:profile=live:muxtype=ts", g.Output, g.ID),
	}

	log.Printf("Executing gpac command: %s", strings.Join(args, " "))

	cmd := exec.Command("gpac", args...)

	return cmd.Run()
}

func NewGPACCommand(id, address, output string) CommandExecutor {
	return &GPACCommand{
		ID:      id,
		Address: address,
		Output:  output,
	}
}
