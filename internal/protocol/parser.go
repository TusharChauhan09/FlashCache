// ! parse raw input -> Command
package protocol

import (
	"errors"
	"strings"
)

var ErrEmptyCommand = errors.New("empty command")

func Parse(input string) (*Command, error){
	input = strings.TrimSpace(input)

	if input==""{
		return nil,ErrEmptyCommand
	}

	parts := strings.Fields(input)

	cmd := &Command{
		Name: strings.ToUpper(parts[0]),
	}

	if len(parts)>1{
		cmd.Args = parts[1:]
	}

	return cmd, nil
}