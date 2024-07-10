package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codecrafters-io/redis-starter-go/app/cache"
	"github.com/codecrafters-io/redis-starter-go/app/cli"
	"github.com/codecrafters-io/redis-starter-go/app/config"
	"github.com/codecrafters-io/redis-starter-go/app/redis"
)

const (
	DEFAULT_PORT = 6379
	DEFAULT_HOST = "0.0.0.0"
)

func main() {
	cfg := config.Config{
		Host: DEFAULT_HOST,
		Port: DEFAULT_PORT,
	}

	if err := cli.ParseFlags(os.Args[1:], &cfg); err != nil {
		log.Fatalf("error parsing args: %v", err.Error())
	}

	fmt.Println(cfg.ReplicaOf)

	storage := cache.NewCache()

	server := redis.NewServer(&cfg, storage)

	err := server.Run()
	if err != nil {
		log.Fatalf("error running server: %v", err.Error())
	}
}
