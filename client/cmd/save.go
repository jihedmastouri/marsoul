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
		filePath, err := cmd.Flags().GetString("file-path")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Reading file failed with: ", err)
			os.Exit(1)
		}

		err = pkg.Save(filePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, "File failed to save with:", err)
			os.Exit(1)
		}
	},
}
