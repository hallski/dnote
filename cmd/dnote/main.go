package main

import (
	"log"
	"os"
	"strconv"

	"dnote/mdfiles"
	"dnote/search"

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
	notes, err := mdfiles.Load(getNotesPath())
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

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a note",
	Long:  "Opens note with ID in Vim",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		notes := loadNotes()

		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Error while editing file %v", err)
		}

		Open(id, notes)
	},
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create and open new note",
	Long:  "Creates a new note with the next available ID and opens it in editor",
	Run: func(cmd *cobra.Command, _ []string) {
		notes := loadNotes()

		note, err := notes.CreateNote()
		if err != nil {
			log.Fatalf("Couldn't create new note: %s", err)
		}

		if err := Edit(note); err != nil {
			log.Fatalf("Error opening new note %s", err)
		}
	},
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all notes",
	Long:  "List all files together with ID",
	Run: func(cmd *cobra.Command, args []string) {
		notes := loadNotes()

		List(notes, os.Stdout)
	},
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search note titles",
	Long:  "Search note titles for strings containing query and list as index",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		notes := loadNotes()

		result := search.NewTitleSearch(os.Args[2], notes)
		ListNoteLinks(result, os.Stdout)
	},
}

var idsCmd = &cobra.Command{
	Use:   "ids",
	Short: "Show matching IDs",
	Long:  "Show matching IDs in an index link list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		notes := loadNotes()

		var ids []int
		for _, strId := range os.Args[2:] {
			id, err := strconv.Atoi(strId)
			if err != nil {
				continue
			}

			ids = append(ids, id)
		}

		result := search.NewIdsSearch(ids, notes)
		ListNoteLinks(result, os.Stdout)
	},
}

func main() {
	rootCmd.AddCommand(openCmd)
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(idsCmd)

	rootCmd.Execute()
}
