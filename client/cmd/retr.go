package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var retrCmd = &cobra.Command{
	Use:   "retr",
	Short: "Retireve (download) a file",
	Long: `Used to retrieve a file from the file nodes
	marsoul retr <file-id>
	`,
	Run: retrieveFile,
}

func retrieveFile(cmd *cobra.Command, args []string) {
	fmt.Println("Helllllloooo")
}
