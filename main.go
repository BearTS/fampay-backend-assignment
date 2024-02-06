package main

import (
	"github.com/BearTS/fampay-backend-assignment/cmd"
	"github.com/BearTS/fampay-backend-assignment/pkg/config"
)

func main() {
	config.ReadFromEnv()
	cmd.Execute()
}
