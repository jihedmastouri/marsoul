package pkg

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jihedmastouri/marsoul/client/internal"
	resolver "github.com/jihedmastouri/marsoul/resolver/transport"
)

var lastUsedFileAddr string

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

	resolverClient, err := getResolvers()
	if err != nil {
		log.Fatal(err)
	}

	addr, err := resolverClient.Save(resolver.SaveRqPayload{
		FileName: fileStat.Name(),
		Size:     int(fileStat.Size()),
		Replicas: 3,
		Region:   "TN",
	})
	log.Println(addr)

	return err
}

func getResolvers() (*resolver.DefaultClient, error) {
	certPath, err := filepath.Abs("../resolver/server-cert.pem")
	if err != nil {
		return nil, err
	}

	keyPath, err := filepath.Abs("../resolver/server-key.pem")
	if err != nil {
		return nil, err
	}

	return resolver.NewClient("localhost:4220", certPath, keyPath), nil
}
