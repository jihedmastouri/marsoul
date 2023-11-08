package internal

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

const (
	gb = 1024 * 1024 * 1024
	mb = 1024 * 1024
	kb = 1024
)

const (
	ConfigLocation = ".marsoul"
	ResolversFile  = "resolvers"
)

type fileSizeUnits struct {
	gb int64
	mb int64
	kb int64
}

func CreateConfigFile(filename string) (*os.File, error) {
	usr, err := user.Current()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Getting User failed with: ", err)
		os.Exit(1)
	}

	configPath := filepath.Join(usr.HomeDir, ConfigLocation)

	if err = os.Mkdir(configPath, 0775); err != nil && !os.IsExist(err) {
		return nil, err
	}

	file, err := os.Create(filepath.Join(configPath, filename))
	if err != nil && !os.IsExist(err) {
		return nil, err
	}

	return file, nil
}

func PrettyFileSize(fileSize int64) string {
	units := bytesToUnits(fileSize)
	res := ""

	if units.gb > 0 {
		res += fmt.Sprintf("%d Gb ", units.gb)
	}
	if units.mb > 0 {
		res += fmt.Sprintf("%d Mb ", units.mb)
	}
	if units.kb > 0 {
		res += fmt.Sprintf("%d Kb", units.kb)
	}
	if res == "" {
		return "< 1 Kb"
	}

	return res
}

func bytesToUnits(fileSize int64) fileSizeUnits {
	var res fileSizeUnits

	if (fileSize / gb) > 0 {
		res.gb = fileSize / gb
		fileSize %= gb
	}
	if (fileSize / mb) > 0 {
		res.mb = fileSize / mb
		fileSize %= mb
	}
	res.kb = fileSize / kb

	return res
}
