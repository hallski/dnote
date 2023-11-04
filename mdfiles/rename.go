package mdfiles

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

func changeID(mdd *MdDirectory, oldID, newID string) error {
	oldpath := mdd.notePath(oldID)
	newpath := mdd.notePath(newID)

	if err := os.Rename(oldpath, newpath); err != nil {
		return err
	}

	updateLinks(mdd, oldID, newID)
	// Update links
	// Rename file

	return nil
}

func updateLinks(mdd *MdDirectory, oldId, newId string) error {
	return filepath.WalkDir(mdd.Path, func(path string, _ fs.DirEntry, e error) error {
		if filepath.Ext(path) != ".md" {
			return nil
		}

		simple := regexp.MustCompile(fmt.Sprintf("\\[\\[%s\\]\\]", oldId))

		content, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}

		simpleReplacement := []byte(fmt.Sprintf("[[%s]]", newId))

		newContent := simple.ReplaceAll(content, simpleReplacement)

		os.WriteFile(path, newContent, 0644)

		return nil
	})
}
