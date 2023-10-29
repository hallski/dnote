package shell

import (
	"fmt"
	"os"
	"path/filepath"

	"hallski/thinkadus/core"
)

func ReadFiles(path string) ([]core.Note, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	notes := make([]core.Note, len(entries))
	for i, entry := range entries {
		note, err := ReadFile(path, entry.Name())
		if err != nil {
			return nil, err
		}
		notes[i] = note
	}

	return notes, nil
}

func ReadFile(dir, fileName string) (core.Note, error) {
	path := filepath.Join(dir, fileName)
	f, err := os.Open(path)
	if err != nil {
		return core.Note{Name: fileName, Path: path, Size: 0}, err
	}
	defer f.Close()

	thing, err := f.Stat()
	if err != nil {
		return core.Note{Name: fileName, Path: path, Size: 0}, err
	}

	return core.Note{Name: fileName, Path: path, Size: thing.Size()}, nil
}

func ReadNoteContent(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	fmt.Println("Read content", string(bytes))

	return string(bytes), nil
}
