package cmd

import (
	"fmt"
	"log"

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
	Run:       saveFile,
}

func saveFile(cmd *cobra.Command, args []string) {
	file, err := cmd.Flags().GetString("file-path")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)
}
