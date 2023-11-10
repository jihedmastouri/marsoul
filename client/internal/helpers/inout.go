package helpers

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var (
	Cyan *color.Color = color.New(color.FgCyan)
	Red               = color.New(color.FgRed)
)

func ReadStdin(prompt string) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	Cyan.Println()
	var res string

	for scanner.Scan() {
		res = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return res, nil
}

func ErrExit(exitMsg string, err error) {
	Cyan.Fprintln(os.Stderr, exitMsg, err)
	os.Exit(1)
}
