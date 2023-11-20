package mdfiles

import "os"

func AddToInbox(path, content string) error {
	toAdd := "- [ ] " + content

	return AddToFile(path, toAdd)
}

func AddToFile(path string, content string) error {
	oldContent, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	tmpPath := path + "~"
	tmpFile, err := os.OpenFile(tmpPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	if _, err := tmpFile.Write([]byte(oldContent)); err != nil {
		return err
	}

	if _, err := tmpFile.Write([]byte("\n" + content)); err != nil {
		return err
	}

	if err := tmpFile.Close(); err != nil {
		return err
	}

	if err := os.Remove(path); err != nil {
		return err
	}

	return os.Rename(tmpPath, path)
}
