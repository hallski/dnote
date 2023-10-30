package main

import (
	"testing"
	"time"
)

// Newer notes have IDs of the form YYMMddHHmmss
func TestParseZkID(t *testing.T) {
	got, err := ParseZkID("221219154814")
	if err != nil {
		t.Errorf("Error in parse function: %s", err)
	}

	expected, _ := time.Parse("2006-01-02 15:04:05", "2022-12-19 15:48:14")
	if got != expected {
		t.Errorf("parsed date didn't match expectation %v, got %v", expected, got)
	}
}

// Older notes have an ID with 4 digits for year and no seconds
func TestOldZkId(t *testing.T) {
	got, err := ParseZkID("202111172324")
	if err != nil {
		t.Errorf("Error in parse function: %s", err)
	}

	expected, _ := time.Parse("2006-01-02 15:04:05", "2021-11-17 23:24:00")
	if got != expected {
		t.Errorf("parsed date didn't match expectation %v, got %v", expected, got)
	}
}

// Test dummy ID for really old notes
func TestSpecialZkID(t *testing.T) {
	got, err := ParseZkID("210101000098")
	if err != nil {
		t.Errorf("Error in parse function: %s", err)
	}

	expected, _ := time.Parse("2006-01-02 15:04:05", "2020-01-01 00:00:00")
	if got != expected {
		t.Errorf("parsed date didn't match expectation %v, got %v", expected, got)
	}
}
