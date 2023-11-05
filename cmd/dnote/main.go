package main

import (
	"fmt"
	"log"
	"os"

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

		Open(mdfiles.PadID(args[0]), notes)
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
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
		List(result, os.Stdout)
	},
}

var idsCmd = &cobra.Command{
	Use:   "ids",
	Short: "Show matching IDs",
	Long:  "Show matching IDs in an index link list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		notes := loadNotes()

		var ids []string
		for _, id := range os.Args[2:] {
			ids = append(ids, mdfiles.PadID(id))
		}

		result := search.NewIdsSearch(ids, notes)

		ListNoteLinks(result, os.Stdout)
	},
}

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a note file",
	Long:  "Rename a note file and update all links to it",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		notes := loadNotes()
		if err := notes.Rename(args[0], args[1]); err != nil {
			fmt.Printf("Failed to rename %s to %s: %s\n", args[0], args[1], err)
		}
	},
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate notebook",
	Long:  "Migrate notebook to latest version",
	Run: func(cmd *cobra.Command, args []string) {
		notes := loadNotes()
		if err := notes.Migrate(); err != nil {
			fmt.Printf("Failed to migrate notebook: %s\n", err)
		}
	},
}

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View note",
	Long:  "View a note without opening it in editor",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		notes := loadNotes()
		note := notes.FindNote(mdfiles.PadID(args[0]))
		View(note)
	},
}

func main() {
	rootCmd.AddCommand(openCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(idsCmd)
	rootCmd.AddCommand(renameCmd)
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(viewCmd)

	rootCmd.Execute()
}
