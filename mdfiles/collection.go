package mdfiles

import (
	"os"
	"path"
	"strings"
)

var filename = path.Join(os.Getenv("HOME"), "store.txt")

func SaveToCollection(id string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(" " + id); err != nil {
		return err
	}
	return nil
}

func ResetCollection() error {
	return os.WriteFile(filename, []byte(""), 0644)
}

func GetCollection() ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	ids := strings.Split(string(content), " ")
	return ids, nil
}
