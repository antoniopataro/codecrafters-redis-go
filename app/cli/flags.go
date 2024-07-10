package cli

import (
	"fmt"
	"strconv"

	"github.com/codecrafters-io/redis-starter-go/app/config"
)

func ParseFlags(args []string, cfg *config.Config) error {
	length := len(args)

	i := 0
	for i < length {
		arg := args[i]

		switch arg {
		case "--port":
			if i+1 > length {
				return fmt.Errorf("not enough arguments")
			}

			port, err := strconv.Atoi(args[i+1])
			if err != nil {
				return fmt.Errorf("port must be a number")
			}

			cfg.Port = port

			i += 1

			continue
		case "--replicaof":
			if i+2 > length {
				return fmt.Errorf("not enough arguments")
			}

			host := args[i+1]
			port, err := strconv.Atoi(args[i+2])
			if err != nil {
				return fmt.Errorf("port must be a number")
			}

			cfg.ReplicaOf = &config.Config{
				Host: host,
				Port: port,
			}

			i += 2

			continue
		default:
			i += 1

			continue
		}
	}

	return nil
}
