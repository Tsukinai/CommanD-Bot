package core

import (
	"strings"
)

type ParsedCommand struct {
	Command string
	Option  string
	Args    []string
}

func ParseMessage(message string) (*ParsedCommand, error) {
	inputArgs := strings.Fields(message)
	if len(inputArgs) <= 1 {
		return nil, NewError("ParseMessage()", "given arguments is <= 1")
	}

	if !strings.HasPrefix(inputArgs[0], "!") {
		return nil,  nil
	}

	pc := &ParsedCommand{
		Command: inputArgs[0],
		Option:  inputArgs[1],
	}

	if len(inputArgs) > 2 {
		pc.Args = inputArgs[2:]
	}

	return pc, nil
}