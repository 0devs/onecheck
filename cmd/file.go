package cmd

import (
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Check file exists and match conditions",
	Run: func(cmd *cobra.Command, args []string) {
		fs := afero.NewOsFs()
		exists, err := afero.Exists(fs, "./main.go")

		zap.L().Info("test", zap.Bool("exists", exists), zap.Error(err))

		//fmt.Println("file called", exists, err)


	},
}

func init()  {
	//fileCmd.Flags().
	rootCmd.AddCommand(fileCmd)
}