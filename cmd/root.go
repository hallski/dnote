package cmd

import (
	"log"
	"os"

	"dnote/config"
	"dnote/mdfiles"
	"dnote/tui/render"

	"github.com/spf13/cobra"
)

var notes *mdfiles.MdDirectory

var title string

func getNotesPath() string {
	path := os.Getenv("DNOTES")
	if path == "" {
		var err error
		path, err = os.Getwd()
		if err != nil {
			path = "."
		}
	}

	return path
}

func loadNotes() *mdfiles.MdDirectory {
	loadedNotes, err := mdfiles.Load(getNotesPath())
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

	dir := getNotesPath()
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
