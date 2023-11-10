package internal_test

import (
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/jihedmastouri/marsoul/client/internal"
)

func TestCreateConfigFile(t *testing.T) {
	filename := "file_test.test_" + strconv.Itoa(rand.Int())

	configDir := internal.GetConfigDir()
	path := filepath.Join(configDir, filename)

	t.Log(path)
	t.Cleanup(func() {
		os.Remove(path)
	})

	err := internal.CreateConfigFile(filename)
	if err != nil {
		if os.IsExist(err) {
			t.Error("file doesn't exits but got this error: ", err)
		}
		t.Error("could not create file: ", err)
	}

	file, err := os.OpenFile(path, os.O_WRONLY, 0644)
	if os.IsExist(err) {
		t.Error("opening file failed with: ", err)
	}
	defer file.Close()

	content := "foo bar baz"
	_, err = file.WriteString(content)
	if err != nil {
		t.Error("could not write to file: ", err)
	}

	buf, err := os.ReadFile(path)
	if err != nil {
		t.Error("could not read from file: ", err)
	}
	if content != string(buf) {
		t.Errorf("content read is not the same. content: %s, buf: %s", content, buf)
	}

	err = internal.CreateConfigFile(filename)
	if err == nil {
		t.Error("file should not have been re-created")
	}
	if err != nil && !os.IsExist(err) {
		t.Error("error should be `file already exists`. Instead got: ", err)
	}
}
