package cmd

import (
	"fmt"
	"os"

	"github.com/jihedmastouri/marsoul/client/pkg"
	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save (upload) a file",
	Long: `Save a file by uploading it into the system
	marsoul save <file>
	`,
	Example:   "",
	ValidArgs: []string{},
	Version:   "v0.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "No file Provided")
			os.Exit(1)
		}

		for _, arg := range args {
			go func(filepath string) {
				if err := pkg.Save(filepath); err != nil {
					fmt.Fprintf(os.Stderr, "File `%s` upload failed with: %s", filepath, err)
				}
			}(arg)
		}

	},
}
