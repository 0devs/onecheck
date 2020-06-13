package filePlugin

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var Name string = "file"
var Version string = "0.0.0"

type InitFn func(*cobra.Command)

func Init(rootCmd *cobra.Command) {
	logger := zap.L()
	logger.Debug("init plugin", zap.String("name", Name), zap.String("version", Version))
	rootCmd.AddCommand(command)
}
