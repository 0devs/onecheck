package main

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var Name string = "file"

var fileCmd = &cobra.Command{
	Use:   Name,
	Short: "Check file exists and match conditions",
	Run: func(cmd *cobra.Command, args []string) {
		logger := zap.L()
		fs := afero.NewOsFs()
		exists, err := afero.Exists(fs, "./main.go")

		logger.Info("test plugin", zap.Bool("exists", exists), zap.Error(err))
	},
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(fileCmd)
}
