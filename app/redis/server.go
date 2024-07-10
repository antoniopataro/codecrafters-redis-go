package redis

import (
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/codecrafters-io/redis-starter-go/app/cache"
	"github.com/codecrafters-io/redis-starter-go/app/command"
	"github.com/codecrafters-io/redis-starter-go/app/config"
	"github.com/codecrafters-io/redis-starter-go/app/input"
)

const (
	DEFAULT_BUFFER_SIZE = 512
)

type Server struct {
	config  *config.Config
	storage *cache.Cache
}

func (s *Server) Run() error {
	address := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to bind to port %d\n", s.config.Port)
	}
	defer listener.Close()

	connections := make(chan net.Conn)
	defer close(connections)

	go s.acceptConnections(
		listener,
		connections,
	)

	for connection := range connections {
		go s.handleConnection(connection)
	}

	return nil
}

func (s *Server) acceptConnections(listener net.Listener, connections chan<- net.Conn) {
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting connection: ", err.Error())
			continue
		}

		connections <- connection
	}
}

func (s *Server) handleConnection(connection net.Conn) {
	defer connection.Close()

	handler := command.NewHandler(s.config, s.storage)

	for {
		var buffer [DEFAULT_BUFFER_SIZE]byte
		writtenBytesLength, err := connection.Read(buffer[:])
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			fmt.Println("error reading: ", err.Error())
			break
		}

		input, err := input.NewInput(buffer[:writtenBytesLength])
		if err != nil {
			fmt.Println("error reading input: ", err.Error())
		}

		output, err := handler.HandleCommand(input)
		if err != nil {
			fmt.Println("error handling command: ", err.Error())
			break
		}
		if output == nil {
			break
		}

		if _, err := connection.Write(output); err != nil {
			fmt.Println("error writing output: ", err.Error())
			break
		}
	}
}

func NewServer(config *config.Config, storage *cache.Cache) *Server {
	return &Server{
		config:  config,
		storage: storage,
	}
}
