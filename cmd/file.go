package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Check file exists and match conditions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("file called")
	},
}

func init()  {
	rootCmd.AddCommand(fileCmd)
}