package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"dnote/config"
	"dnote/mdfiles"
	"dnote/tui/render"

	"github.com/spf13/cobra"
)

var notes *mdfiles.MdDirectory

var title string

func expandPath(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("Failed to expand user homedir")
		}

		return filepath.Join(homeDir, path[2:]), nil
	}

	return path, nil
}

func getNotesPath() (string, error) {
	cfg := config.GetConfig()

	possiblePaths := []string{os.Getenv("DNOTES"), cfg.Notes.DefaultDir}

	for _, path := range possiblePaths {
		if path == "" {
			continue
		}

		expandedPath, err := expandPath(path)
		if err != nil {
			return "", err
		}
		return expandedPath, nil
	}

	return "", fmt.Errorf("No paths found")
}

func loadNotes() *mdfiles.MdDirectory {
	path, err := getNotesPath()
	if err != nil {
		panic(err)
	}
	loadedNotes, err := mdfiles.Load(path)
	if err != nil {
		panic(err)
	}

	return loadedNotes
}

var rootCmd = &cobra.Command{
	Use:     "dnote",
	Short:   "dNote system",
	Long:    "My personal note system",
	Version: "0.1",
}

func Execute() {
	// Initialize configuration
	if err := config.InitConfig(); err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}

	// Initialize styles after config is loaded
	render.InitializeStyles()

	dir, err := getNotesPath()
	if err != nil {
		log.Fatalf("Could not find notes path, set DNOTES or configure it in config")
		return
	}

	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.Mkdir(dir, 0700)
	} else {
		if !info.IsDir() {
			log.Fatalf("%s is not a directory", dir)
			return
		}
	}

	os.Chdir(dir)

	notes = loadNotes()
	if notes.IsEmpty() {
		bootstrapDirectory()
	}
	notes = loadNotes()

	rootCmd.Execute()
}
