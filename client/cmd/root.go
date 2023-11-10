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
	initCmd.Flags().StringP("file-path", "f", "", "copy address from a file")
}
