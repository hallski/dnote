package cmd

import (
	"os"

	"dnote/mdfiles"

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
	var err error
	notes, err = mdfiles.Load(getNotesPath())
	if err != nil {
		panic(err)
	}

	return notes
}

var rootCmd = &cobra.Command{
	Use:     "dnote",
	Short:   "dNote system",
	Long:    "My personal note system",
	Version: "0.1",
}

func Execute() {
	dir := getNotesPath()
	os.Chdir(dir)

	loadNotes()

	rootCmd.Execute()
}
