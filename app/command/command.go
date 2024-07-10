package command

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/app/cache"
	"github.com/codecrafters-io/redis-starter-go/app/config"
	"github.com/codecrafters-io/redis-starter-go/app/input"
)

type Arg string

const (
	commandArg Arg = "command"
	echoArg    Arg = "echo"
	getArg     Arg = "get"
	infoArg    Arg = "info"
	pingArg    Arg = "ping"
	setArg     Arg = "set"
	pxArg      Arg = "px"
)

const (
	errUnknownCommandResponse = "-ERR unknown command\r\n"
	nullBulkCommandResponse   = "$-1\r\n"
	okCommandResponse         = "+OK\r\n"
	pongResponse              = "+PONG\r\n"
)

type Handler struct {
	serverConfig  *config.Config
	sharedStorage *cache.Cache
}

func (h *Handler) HandleCommand(input *input.Input) ([]byte, error) {
	nextArg, err := input.ParseNextArg()
	if err != nil {
		return nil, fmt.Errorf("could not parse next arg: %v", err.Error())
	}

	command := Arg(strings.ToLower(string(nextArg.Value)))
	switch command {
	case commandArg:
		return h.handleCommandCommand()
	case echoArg:
		return h.handleEchoCommand(input)
	case getArg:
		return h.handleGetCommand(input)
	case infoArg:
		return h.handleInfoCommand(input)
	case pingArg:
		return h.handlePingCommand()
	case setArg:
		return h.handleSetCommand(input)
	default:
		return h.handleUnknownCommand()
	}
}

func NewHandler(serverConfig *config.Config, sharedStorage *cache.Cache) *Handler {
	return &Handler{
		serverConfig:  serverConfig,
		sharedStorage: sharedStorage,
	}
}
