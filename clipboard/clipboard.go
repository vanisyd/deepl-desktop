package clipboard

import (
	"errors"
	"os/exec"
)

const noAvailableCMDError = "no commands available on your system"

type Clipboard struct {
	Commands    []Cmd
	MainCommand *Cmd
}

type Cmd struct {
	Command string
	Params  []string
}

var AvailableCommands []Cmd = []Cmd{
	{
		Command: "xsel",
		Params:  []string{"-ob"},
	},
	{
		Command: "xclip",
		Params:  []string{"-o"},
	},
}

func (c *Clipboard) Init() error {
	for _, command := range AvailableCommands {
		if _, err := exec.LookPath(command.Command); err == nil {
			c.MainCommand = &command
			break
		}
	}

	if c.MainCommand == nil {
		return errors.New(noAvailableCMDError)
	}

	return nil
}

func (c *Clipboard) Read() (string, error) {
	if c.MainCommand == nil {
		return "", errors.New(noAvailableCMDError)
	}

	cmd := exec.Command(c.MainCommand.Command, c.MainCommand.Params...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
