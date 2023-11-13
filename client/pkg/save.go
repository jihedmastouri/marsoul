package pkg

import (
	"fmt"
	"net"
	"os"

	"github.com/jihedmastouri/marsoul/client/internal"
)

var lastUsedResolverAddr string

func Save(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	fileStat, err := file.Stat()
	if err != nil {
		return err
	}

	switch {
	case fileStat.IsDir():
		return fmt.Errorf("File is directory")
	case fileStat.Size() > internal.MaxSizeFile:
		return fmt.Errorf("File size limit")
	}

	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		return err
	}
	defer l.Close()

	fmt.Println(internal.PrettyFileSize(fileStat.Size()))

	return err
}

func getResolvers() {

}
