package main

import (
	"github.com/0devs/onecheck/cmd"
	"github.com/0devs/onecheck/logger"
)

func main() {
	logger.InitLogger()
	cmd.Execute()
}
