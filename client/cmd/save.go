package cmd

import (
	"fmt"
	"sync"

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

		wg := sync.WaitGroup{}
		for _, arg := range args {
			wg.Add(1)
			go func(filepath string) {
				if err := pkg.Save(filepath); err != nil {
					errStr := fmt.Sprintf("File `%s` upload failed with: ", filepath)
					helpers.ErrExit(errStr, err)
				}
				wg.Done()
			}(arg)
		}
		wg.Wait()
	},
}
