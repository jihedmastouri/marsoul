package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jihedmastouri/marsoul/client/internal"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup resolvers as a file under $home/.marsoul",
	Long: `Create a file under $home/.marsoul with a list of resolvers passed as args
	marsoul init <address...>
	`,
	Example: "",
	Run: func(cmd *cobra.Command, args []string) {
		filePath, err := cmd.Flags().GetString("file-path")

		if len(args) == 0 && err != nil {
			if len(args) == 0 {
				fmt.Fprintln(os.Stderr, "No resolvers passed as args or as file")
				os.Exit(1)
			}
			fmt.Fprintln(os.Stderr, "Reading file-path failed with: ", err)
			os.Exit(1)
		}

		if filePath != "" {
			providedFile, err := os.Open(filePath)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Opening file failed with: ", err)
				os.Exit(1)
			}
			defer providedFile.Close()
			scanner := bufio.NewScanner(providedFile)
			for scanner.Scan() {
				line := scanner.Text()
				args = append(args, line)
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Reading file content failed with: ", err)
				os.Exit(1)
			}

		}

		if valid, addrs := internal.ValidateAdrs(args...); !valid {
			fmt.Fprintln(os.Stderr, "Argument passed is not a valid address: ", addrs)
			os.Exit(1)
		}

		file, err := internal.CreateConfigFile(internal.ResolversFile)

		writable := strings.Join(args, "\n")
		_, err = file.WriteString(writable)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Writing to file `$HOME/.marsoul/resolvers` failed with: ", err)
			os.Exit(1)
		}
	},
}
