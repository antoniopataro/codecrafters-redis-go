package command

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/app/input"
	"github.com/codecrafters-io/redis-starter-go/app/resp"
)

func (h *Handler) handleGetCommand(input *input.Input) ([]byte, error) {
	arg1, err := input.ParseNextArg()
	if err != nil {
		return nil, fmt.Errorf("could not parse next arg: %v", err.Error())
	}

	getKey := string(arg1.Value)

	value, ok := h.sharedStorage.Get(getKey)
	if !ok {
		return []byte(nullBulkCommandResponse), nil
	}

	valueStr, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("could not cast value to string")
	}

	return resp.SimpleString(valueStr), nil
}
