package parser

import (
	"errors"
	"strings"
)

type Command struct {
	Action, Key, Value string
}

func Parse(input string) (Command, error) {
	inputParts := strings.Fields(input)
	partsLen := len(inputParts)
	if partsLen == 0 {
		return Command{}, errors.New("empty command")
	}

	action := strings.ToUpper(inputParts[0])
	if (action == "GET" && partsLen < 2) || (action == "SET" && partsLen < 3) {
		return Command{}, errors.New("not enough parameters")
	}
	if partsLen > 3 && !(strings.HasPrefix(inputParts[2], "\"") && strings.HasSuffix(inputParts[partsLen - 1], "\"")){
		return Command{}, errors.New("too many parameters")	
	}

	key := inputParts[1]
	value := strings.Join(inputParts[2:], " ")
	normalizedValue := strings.Trim(value, "\"")

	return Command{action, key, normalizedValue}, nil
}