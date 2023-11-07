package main

import (
	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:       "marsoul save <file>",
	Short:     "Save (upload) a file",
	Long:      "Save a file by uploading it into the system",
	Example:   "",
	ValidArgs: []string{},
	Version:   "v0.0.0",
	Run:       func(cmd *cobra.Command, args []string) {},
}
