package filePlugin

import "github.com/spf13/cobra"

//"github.com/spf13/afero"

var command = &cobra.Command{
	Use:   Name,
	Short: "Check file exists and match conditions",
	Run: func(cmd *cobra.Command, args []string) {
		//fs := afero.NewOsFs()
		//exists, err := afero.Exists(fs, "./main.go")
	},
}
