package main

import (
	"github.com/spf13/cobra"
)

var retrCmd = &cobra.Command{
	Use:   "marsoul retrieve <file-id>",
	Short: "Retireve (download) a file",
	Long:  "Used to retrieve a file from the file nodes",
	Run:   func(cmd *cobra.Command, args []string) {},
}
