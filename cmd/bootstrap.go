package cmd

import (
	"dnote/mdfiles"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Bootstrap an empty notes directory",
	Long:  "Initialize an empty directory with inbox (000.md) and index (001.md) files",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := bootstrapDirectory()
		if err != nil {
			fmt.Printf("Error bootstrapping directory: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Successfully bootstrapped notes directory with inbox and index files")
	},
}

func init() {
	rootCmd.AddCommand(bootstrapCmd)
}

func bootstrapDirectory() (*mdfiles.MdDirectory, error) {
	dir := getNotesPath()

	// Check if directory already has markdown files
	files, err := filepath.Glob(filepath.Join(dir, "*.md"))
	if err != nil {
		return nil, fmt.Errorf("failed to check for existing files: %v", err)
	}

	if len(files) > 0 {
		return nil, fmt.Errorf("directory already contains markdown files, cannot bootstrap")
	}

	currentTime := time.Now().Format("2006-01-02 15:04")

	// Create inbox file (000.md)
	inboxContent := fmt.Sprintf("[:created]: _ \"%s\"\n\n# Inbox\n", currentTime)

	err = os.WriteFile(filepath.Join(dir, "000.md"), []byte(inboxContent), 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to create inbox file: %v", err)
	}

	const editShortcut = "`e`"
	const addShortcut = "`a`"
	const configFile = "`dnote.yaml`"
	// Create index file (001.md)
	indexContent := fmt.Sprintf(`# Index

Welcome to dNote!

Please see the [README](https://github.com/hallski/dnote/blob/main/README.md) for more documentation.

Otherwise, start by editing this file with %s or add a new file with %s. Make sure you first have followed the instructions and setup your %s configuration file.currentTime

## Inbox

The inbox lives in a file called [[000]].

[:created]: _ "%s"

`, editShortcut, addShortcut, configFile, currentTime)

	err = os.WriteFile(filepath.Join(dir, "001.md"), []byte(indexContent), 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to create index file: %v", err)
	}

	return loadNotes(), nil
}
