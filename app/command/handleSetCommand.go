package command

import (
	"fmt"
	"strconv"

	"github.com/codecrafters-io/redis-starter-go/app/cache"
	"github.com/codecrafters-io/redis-starter-go/app/input"
)

func (h *Handler) handleSetCommand(input *input.Input) ([]byte, error) {
	if input.Length < 2 {
		return nil, fmt.Errorf("invalid number of arguments")
	}

	arg1, err := input.ParseNextArg()
	if err != nil {
		return nil, fmt.Errorf("could not parse next arg: %v", err.Error())
	}
	arg2, err := input.ParseNextArg()
	if err != nil {
		return nil, fmt.Errorf("could not parse next arg: %v", err.Error())
	}

	var expiration *int

	if input.Length == 2 {
		arg3, err := input.ParseNextArg()
		if err != nil {
			return nil, fmt.Errorf("could not parse next arg: %v", err.Error())
		}

		arg3Arg := Arg(arg3.Value)

		if arg3Arg == pxArg {
			arg4, err := input.ParseNextArg()
			if err != nil {
				return nil, fmt.Errorf("could not parse next arg: %v", err.Error())
			}

			n, err := strconv.Atoi(string(arg4.Value))
			if err != nil {
				return nil, fmt.Errorf("could not convert arg to integer: %v", err.Error())
			}

			expiration = &n
		}
	}

	setKey := string(arg1.Value)
	setValue := string(arg2.Value)

	record := cache.Record{
		Expiration: expiration,
		Value:      setValue,
	}
	h.sharedStorage.Store(setKey, record)

	return []byte(okCommandResponse), nil
}
