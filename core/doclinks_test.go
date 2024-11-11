package core

import "testing"

func TestStartValue(t *testing.T) {
	l := NewDocLinks([]string{"l1", "l2"})

	if l.Current() != "" {
		t.Errorf("expected empty string, got %s", l.Current())
	}
}

func TestEmpty(t *testing.T) {
	l := NewDocLinks([]string{})

	if l.Current() != "" {
		t.Errorf("expected empty string, got %s", l.Current())
	}
}

func TestNextPrev(t *testing.T) {
	l := NewDocLinks([]string{"l1", "l2", "l3"})

	l.Next()
	if l.Current() != "l1" {
		t.Errorf("expected l1, got %s", l.Current())
	}
	l.Next()
	if l.Current() != "l2" {
		t.Errorf("expected l2, got %s", l.Current())
	}
	l.Prev()
	if l.Current() != "l1" {
		t.Errorf("expected l1, got %s", l.Current())
	}
}

func TestWrapping(t *testing.T) {
	l := NewDocLinks([]string{"l1", "l2", "l3"})

	l.Prev()
	if l.Current() != "l3" {
		t.Errorf("expected l3, got %s", l.Current())
	}

	l.Next()
	if l.Current() != "l1" {
		t.Errorf("expected l1, got %s", l.Current())
	}
}

func TestAssignShortcuts(t *testing.T) {
	l := NewDocLinks([]string{"l1", "l2", "l3"})

	sc := l.GetShortcut("l1")
	if sc != "A" {
		t.Errorf("expected A, got %s", sc)
	}

	sc = l.GetShortcut("l4")
	if sc != "" {
		t.Errorf("expected empty string, got %s", sc)
	}
}

func TestLookupShortcuts(t *testing.T) {
	l := NewDocLinks([]string{"l1", "l2", "l3"})

	sc := l.GetShortcut("l2")
	if sc != "B" {
		t.Errorf("expected l2, got %s", l.GetLink(sc))
	}

	expected := ShortcutLink{}
	if l.GetLinkFromShortcut("D") != expected {
		t.Errorf("expected empty link, got %+v", l.GetLinkFromShortcut("D"))
	}
}

func TestIsActive(t *testing.T) {
	l := NewDocLinks([]string{"l1", "l2", "l3"})

	l.Next()
	l.Next()

	if l.IsActive(1) != true {
		t.Errorf("expected true, got %v", l.IsActive(1))
	}
}
