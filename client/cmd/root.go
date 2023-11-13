package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "marsoul",
	Short: "Useless file transfer system.",
	Long: `A file transer system similar to QbitTorrent
	marsoul [save <file> | retr <file-id>]
	`,
	Example:   "",
	ValidArgs: []string{},
	Version:   "v0.0.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(saveCmd, retrCmd, initCmd)

	//Init
	initCmd.Flags().StringP("file-path", "f", "", "specify a file to copy known resolvers address from.")

	//Save
	saveCmd.Flags().StringSliceP("resolvers-addrs", "a", []string{}, "manually specify a list of addresses for resolvers")
	saveCmd.Flags().StringP("resolver-name", "n", "", "specify the name of resolver (in the know resolvers list) to use.")
	saveCmd.Flags().BoolP("ignore-know-list", "i", false, "Stop the system from falling back to the resolvers know list if a resolver address or name was provided")
}
