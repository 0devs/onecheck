package main

import (
	"github.com/0devs/onecheck/cmd"
	"github.com/0devs/onecheck/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"plugin"
)

func main() {
	logger.InitLogger()
	logger := zap.L()

	plugin, err := plugin.Open("file.so")

	if err != nil {
		logger.Error("error", zap.Error(err))
	}

	v, err := plugin.Lookup("Name")

	if err != nil {
		logger.Error("error", zap.Error(err))
	}

	logger.Info("v = ", zap.String("v", *v.(*string)))

	initPlugin, err := plugin.Lookup("Init")

	if err != nil {
		logger.Error("error", zap.Error(err))
	}

	initPlugin.(func(rootCmd *cobra.Command))(cmd.RootCmd)

	cmd.Execute()
}
