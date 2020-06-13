package main

import (
	_logger "github.com/0devs/onecheck/logger"
	filePlugin "github.com/0devs/onecheck/plugin/file"
	"go.uber.org/zap"
)

var version string = "0.0.0"

func main() {
	_logger.InitLogger()

	logger := zap.L()

	logger.Debug("init onecheck", zap.String("version", version))

	filePlugin.Init(RootCmd)

	Execute()
}
