package internal_test

import (
	"math/rand"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/jihedmastouri/marsoul/client/internal"
)

func TestCreateConfigFile(t *testing.T) {
	filename := "file_test.test_" + strconv.Itoa(rand.Int())

	usr, err := user.Current()
	if err != nil {
		t.Error("Couldn't read file: ", err)
	}
	path := filepath.Join(usr.HomeDir, internal.ConfigLocation, filename)

	t.Log(path)
	t.Cleanup(func() {
		os.Remove(path)
	})

	file, err := internal.CreateConfigFile(filename)
	if err != nil {
		if os.IsExist(err) {
			t.Error("file doesn't exits but got this error: ", err)
		}
		t.Error("could not create file: ", err)
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
	if strings.Compare(content, string(buf)) != 0 {
		t.Errorf("Content read is not the same. content: %s, buf: %s", content, buf)
	}

	f, err := internal.CreateConfigFile(filename)
	if f != nil {
		t.Error("file should not have been re-created")
	}
	if err != nil && !os.IsExist(err) {
		t.Error("error should be `file already exists`. Instead got: ", err)
	}
}
