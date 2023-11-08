package pkg

import (
	"fmt"
	"os"

	"github.com/jihedmastouri/marsoul/client/internal"
)

func Save(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	fileStat, err := file.Stat()
	if err != nil {
		return err
	}

	fmt.Println(internal.PrettyFileSize(fileStat.Size()))

	return err
}
