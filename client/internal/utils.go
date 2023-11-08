package internal

import (
	"fmt"
)

const (
	gb = 1024 * 1024 * 1024
	mb = 1024 * 1024
	kb = 1024
)

type fileSizeUnits struct {
	gb int64
	mb int64
	kb int64
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
