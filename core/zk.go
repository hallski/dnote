package core

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

func ParseZkID(id string) (time.Time, error) {
	const oldFormat = "200601021504"
	const newFormat = "060102150405"
	const specFormat = "0601021504"
	t, err := time.Parse(newFormat, id)
	if err != nil {
		t, err := time.Parse(oldFormat, id)
		if err != nil {
			_, err := time.Parse(specFormat, id[0:len(id)-2])
			if err != nil {
				newError := fmt.Errorf("Failed to parse ZK ID %s: %s", id, err)
				return time.Time{}, newError
			}
			return time.Parse(oldFormat, "202001010000")
		}

		return t, nil
	}

	return t, nil
}

func UpdateLinks(directory, oldId, newId, title string) error {
	return filepath.WalkDir(directory, func(path string, _ fs.DirEntry, e error) error {
		if filepath.Ext(path) != ".md" {
			return nil
		}

		simple := regexp.MustCompile(fmt.Sprintf("\\[\\[%s\\]\\]", oldId))
		extended := regexp.MustCompile(fmt.Sprintf("\\[\\[%s[a-zA-Z\\- ]+\\]\\]", oldId))

		content, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}

		simpleReplacement := []byte(fmt.Sprintf("[[%s]]", newId))
		extReplacement := []byte(fmt.Sprintf("%s [[%s]]", title, newId))

		newContent := simple.ReplaceAll(content, simpleReplacement)
		newContent = extended.ReplaceAll(newContent, extReplacement)

		os.WriteFile(path, newContent, 0644)

		return nil
	})
}
