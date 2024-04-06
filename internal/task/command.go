package task

import "os/exec"

type CommandExecutor interface {
	Execute() error
}

type GPACCommand struct {
	Address string
	Output  string
}

func (g *GPACCommand) Execute() error {
	args := []string{
		"-i", g.Address,
		"@",
		"c=avc:b=1m",
		"@@", "c=avc:b=2m",
		"@@", "c=aac:b=44k",
		"@",
		"@1", "@2",
		"-o", g.Output,
	}
	cmd := exec.Command("gpac", args...)
	return cmd.Run()
}

func NewGPACCommand(address, output string) CommandExecutor {
	return &GPACCommand{
		Address: address,
		Output:  output,
	}
}
