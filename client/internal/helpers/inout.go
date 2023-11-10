package helpers

import (
	"bufio"
	"os"

	"github.com/fatih/color"
)

var (
	Cyan *color.Color = color.New(color.FgCyan)
	Red               = color.New(color.FgRed)
)

func ReadStdin(prompt string) (string, error) {
	Cyan.Print(prompt)

	scanner := bufio.NewScanner(os.Stdin)
	var res string

	for scanner.Scan() {
		res = scanner.Text()
		break
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return res, nil
}

func ErrExit(exitMsg string, err error) {
	Red.Fprintln(os.Stderr, exitMsg, err)
	os.Exit(1)
}
