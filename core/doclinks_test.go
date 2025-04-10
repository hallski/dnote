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
	l := NewDocLinks([]string{"l1", "l2", "l3", "l4", "l5"})

	sc := l.GetShortcut("l1")
	if sc != "A" {
		t.Errorf("expected A, got %s", sc)
	}

	sc = l.GetShortcut("l2")
	if sc != "B" {
		t.Errorf("expected A, got %s", sc)
	}

	sc = l.GetShortcut("l3")
	if sc != "C" {
		t.Errorf("expected A, got %s", sc)
	}

	sc = l.GetShortcut("l4")
	if sc != "D" {
		t.Errorf("expected A, got %s", sc)
	}

	sc = l.GetShortcut("l6")
	if sc != "" {
		t.Errorf("expected empty string, got %s", sc)
	}
}

func TestLookupShortcuts(t *testing.T) {
	l := NewDocLinks([]string{"l1", "l2", "l3", "l4", "l5"})

	link := l.GetLink("l1")
	if link.Shortcut != "A" {
		t.Errorf("expected A, got %s", link.Shortcut)
	}

	link = l.GetLink("l2")
	if link.Shortcut != "B" {
		t.Errorf("expected B, got %s", link.Shortcut)
	}

	link = l.GetLink("l3")
	if link.Shortcut != "C" {
		t.Errorf("expected C, got %s", link.Shortcut)
	}

	link = l.GetLink("l4")
	if link.Shortcut != "D" {
		t.Errorf("expected D, got %s", link.Shortcut)
	}

	link = l.GetLink("l5")
	if link.Shortcut != "E" {
		t.Errorf("expected E, got %s", link.Shortcut)
	}

	expected := ShortcutLink{}
	if l.GetLinkFromShortcut("F") != expected {
		t.Errorf("expected empty link, got %+v", l.GetLinkFromShortcut("F"))
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
