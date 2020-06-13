package command

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type InitFn func(*cobra.Command)

func Init(rootCmd *cobra.Command) {
	logger := zap.L()
	err := initFlags(&commandParams)

	if err != nil {
		logger.Error("error on initFlags", zap.Error(err))
		panic(255)
	}

	rootCmd.AddCommand(command)
}
