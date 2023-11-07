package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:       "marsoul [save <file> | retrieve <file-id>]",
	Short:     "Useless file transfer system.",
	Long:      "A file transer system similar to QbitTorrent",
	Example:   "",
	ValidArgs: []string{},
	Version:   "v0.0.0",
	Run:       func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Init() {
	rootCmd.AddCommand(saveCmd, retrCmd)
}
