package mdfiles

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"dnote/core"
)

func loadNote(path string) (*core.Note, error) {
	// Read a note
	id, err := getFileID(path)
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	sContent := string(content)
	noteDate, err := getDate(sContent)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse date in note, ignoring")
	}

	note := &core.Note{
		ID:      id,
		Path:    path,
		Content: sContent,
		Date:    noteDate,
		Title:   getTitle(sContent),
		Tags:    getTags(sContent),
		Links:   core.ExtractLinks(sContent),
	}

	return note, nil
}

func createNote(path string, id string, title string) (*core.Note, error) {
	var out bytes.Buffer

	// Replace with a template

	fmt.Fprintf(&out, "# %s\n", title)

	fmt.Fprintf(&out, "\n\n[:created]: _ \"%s\"\n",
		time.Now().Format("2006-01-02 15:04"))

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("Failed to create new note: %s", err)
	}
	f.WriteString(out.String())
	defer f.Close()

	note := &core.Note{
		ID:      id,
		Path:    path,
		Content: out.String(),
	}

	return note, nil
}

func getTitle(content string) string {
	re := regexp.MustCompile("# (.*)")

	result := re.FindStringSubmatch(content)
	if result != nil {
		return string(result[1])
	}

	return ""
}

func getDate(content string) (time.Time, error) {
	re := regexp.MustCompile("\\[:created\\]: _ \"([0-9]{4}-[0-9]{2}-[0-9]{2}).*\"")

	result := re.FindStringSubmatch(content)
	if result == nil {
		return time.Now(), fmt.Errorf("Failed to find date")
	}

	date, err := time.Parse("2006-01-02", result[1])
	if err != nil {
		return time.Now(), fmt.Errorf("Failed to parse date: %s", result[1])
	}

	return date, nil
}

func getTags(content string) []string {
	re := regexp.MustCompile(" (#[a-zA-Z-]+)")

	var tags []string
	matches := re.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		tags = append(tags, match[1])
	}

	return tags
}

func getFileID(path string) (string, error) {
	base := filepath.Base(path)

	fileWithoutExt, ext, found := strings.Cut(base, ".")
	if !found || ext != "md" {
		return "", fmt.Errorf("Filename not following convention of id.md: %s",
			path)
	}

	return fileWithoutExt, nil
}

func addTimestampToNote(path string, timestamp time.Time) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("Failed to append timestamp: %s", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	fmt.Fprintf(w, "\n\n[:created]: _ \"%s\"\n",
		timestamp.Format("2006-01-02 15:04"))

	return nil
}

func PadID(id string) string {
	if len(id) >= core.IDLength {
		return id
	}

	zPad := strings.Repeat("0", core.IDLength-len(id))
	return zPad + id
}
