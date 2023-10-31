package mdfiles

import "testing"

func TestGetTitle(t *testing.T) {
	content := `This is a file
# This is a title

This is more content
`

	got := getTitle(content)
	expected := "This is a title"

	if got != expected {
		t.Errorf("title did not match. expected %s, got %s", expected, got)
	}
}
