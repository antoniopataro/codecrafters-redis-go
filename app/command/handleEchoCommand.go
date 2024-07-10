package command

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/app/input"
	"github.com/codecrafters-io/redis-starter-go/app/resp"
)

func (h *Handler) handleEchoCommand(input *input.Input) ([]byte, error) {
	arg1, err := input.ParseNextArg()
	if err != nil {
		return nil, fmt.Errorf("could not parse next arg: %v", err.Error())
	}

	echoContent := string(arg1.Value)

	return resp.SimpleString(echoContent), nil
}
