package command

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/app/input"
	"github.com/codecrafters-io/redis-starter-go/app/resp"
)

func (h *Handler) handleInfoCommand(input *input.Input) ([]byte, error) {
	arg1, err := input.ParseNextArg()
	if err != nil {
		return nil, fmt.Errorf("could not parse next arg: %v", err.Error())
	}

	infoKey := string(arg1.Value)

	switch infoKey {
	case "replication":
		fmt.Println(h.serverConfig.ReplicaOf)
		if h.serverConfig.ReplicaOf != nil {
			return resp.BulkString("role:slave"), nil
		}

		return resp.BulkString("role:master"), nil
	default:
		return nil, nil
	}
}
