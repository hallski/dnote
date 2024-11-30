package core

import (
	"regexp"
)

var LinkRegexp = regexp.MustCompile(`\[\[([a-zA-Z0-9\&\?\- ]+)\]\]`)

// Finds all links in a string and returns them as a list
func ExtractLinks(str string) []string {
	var links []string
	matches := LinkRegexp.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		links = append(links, match[1])
	}

	return links
}
