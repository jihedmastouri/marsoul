package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
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
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "No resolvers passed as args")
			os.Exit(1)
		}

		if valid, addrs := internal.ValidateAdrs(args...); !valid {
			fmt.Fprintln(os.Stderr, "Argument passed is not a valid address: ", addrs)
			os.Exit(1)
		}

		usr, err := user.Current()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Getting User failed with: ", err)
			os.Exit(1)
		}

		configPath := filepath.Join(usr.HomeDir, ".marsoul")

		if err = os.Mkdir(configPath, 0775); err != nil && !os.IsExist(err) {
			fmt.Fprintln(os.Stderr, "Creating `$HOME/.marsoul` failed with: ", err)
			os.Exit(1)
		}

		file, err := os.Create(filepath.Join(configPath, "resolvers"))
		if err != nil && !os.IsExist(err) {
			fmt.Fprintln(os.Stderr, "Creating file `$HOME/.marsoul/resolvers` failed with: ", err)
			os.Exit(1)
		}

		writable := strings.Join(args, "\n")
		_, err = file.WriteString(writable)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Writing to file `$HOME/.marsoul/resolvers` failed with: ", err)
			os.Exit(1)
		}
	},
}
