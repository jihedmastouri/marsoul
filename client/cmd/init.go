package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jihedmastouri/marsoul/client/internal"
	"github.com/jihedmastouri/marsoul/client/internal/helpers"
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

		configDir := internal.GetConfigDir()
		absFilePath := filepath.Join(configDir, internal.ResolversFile)

		if len(args) == 0 && err != nil {
			if len(args) == 0 {
				fmt.Fprintln(os.Stderr, "No resolvers passed as args or as file")
				os.Exit(1)
			}
			fmt.Fprintln(os.Stderr, "Reading file-path failed with: ", err)
			os.Exit(1)
		}

		if filePath != "" {
			if err := FileArgs(filePath, args); err != nil {
				fmt.Fprintln(os.Stderr, "Reading file failed with: ", err)
				os.Exit(1)
			}
		}

		if valid, addrs := internal.ValidateAdrs(args); !valid {
			errStr := fmt.Sprintf("Argument passed is not a valid address: %s", addrs)
			helpers.ErrExit(errStr, nil)
		}

		if err = internal.CreateConfigFile(internal.ResolversFile); err != nil {
			if !os.IsExist(err) {
				errStr := fmt.Sprintf("creating file `%s` failed with: ", absFilePath)
				helpers.ErrExit(errStr, err)
			}

			s, err := helpers.ReadStdin("file exists already. Would you like to (A)dd to it or (R)eplace: (a/r) ")
			if err != nil {
				helpers.ErrExit("reading option failed with: ", err)
			}
			if strings.ToLower(s) == "r" {
				os.Truncate(absFilePath, 0)
			}
		}

		file, err := os.OpenFile(absFilePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			errStr := fmt.Sprintf("opening file `%s` failed with: ", absFilePath)
			helpers.ErrExit(errStr, err)
		}
		defer file.Close()

		writable := strings.Join(args, "\n")
		writable += "\n"

		if _, err = file.WriteString(writable); err != nil {
			errStr := fmt.Sprintf("writing to file `%s` failed with: ", absFilePath)
			helpers.ErrExit(errStr, err)
		}
	},
}

func FileArgs(filePath string, args []string) error {
	providedFile, err := os.Open(filePath)
	if err != nil {
		return err

	}
	defer providedFile.Close()

	scanner := bufio.NewScanner(providedFile)
	for scanner.Scan() {
		line := scanner.Text()
		args = append(args, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
