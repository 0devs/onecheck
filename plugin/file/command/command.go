package command

import (
	_result "github.com/0devs/onecheck/internal/result"
	filePlugin "github.com/0devs/onecheck/plugin/file"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

var command = &cobra.Command{
	Use:   filePlugin.Plugin.Name,
	Short: "Check file exists and match conditions",
	Run: func(cmd *cobra.Command, args []string) {
		logger := zap.L()

		logger.Debug("commandParams on run", zap.Any("commandParams", commandParams))

		config := convertCommandParamsToConfig(&commandParams)

		logger.Debug("config on run", zap.Any("config", config))

		err := filePlugin.Plugin.ValidateConfig(&config)

		if err != nil {
			logger.Error("error on validateConfig", zap.Error(err))
			os.Exit(1)
		}

		err, result := filePlugin.Plugin.Run(&config)

		if err != nil {
			logger.Error("error on run", zap.Error(err))
			panic(255)
		}

		logger.Debug("file plugin result", zap.Any("result", result))

		_result.FinishAndExit(&result)
	},
}

func initFlags(commandParams *CommandParams) error {
	logger := zap.L()

	command.Flags().StringVar(&commandParams.Path, "path", "", "Path to file (required)")
	command.Flags().BoolVar(&commandParams.Exists, "exists", true, "File should exists or not")

	if err := command.MarkFlagRequired("path"); err != nil {
		logger.Error("can't mark flag as required", zap.String("flag", "path"))
		return err
	}

	return nil
}
