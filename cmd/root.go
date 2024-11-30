package cmd

import (
	"os"

	"dnote/mdfiles"

	"github.com/spf13/cobra"
)

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
	dir := getNotesPath()
	os.Chdir(dir)

	rootCmd.Execute()
}
