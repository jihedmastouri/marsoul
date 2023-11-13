package cmd

import (
	"fmt"
	"os"

	"github.com/jihedmastouri/marsoul/client/internal/helpers"
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
	PreRun: func(cmd *cobra.Command, args []string) {
		adddr, _ := cmd.Flags().GetStringSlice("resolvers-addrs")
		name, _ := cmd.Flags().GetString("resolver-name")

		ignore, _ := cmd.Flags().GetBool("ignore-know-list")
		if ignore && (len(adddr) == 0 || name == "") {
			helpers.ErrExit("You cannot ignore know resolvers list if you don't provide any address", nil)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			helpers.ErrExit("No file path provided", nil)
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
