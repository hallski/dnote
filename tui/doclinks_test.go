package tui

import "testing"

func TestStartValue(t *testing.T) {
	l := newDocLinks([]string{"l1", "l2"})

	if l.Current() != "" {
		t.Errorf("expected empty string, got %s", l.Current())
	}
}

func TestEmpty(t *testing.T) {
	l := newDocLinks([]string{})

	if l.Current() != "" {
		t.Errorf("expected empty string, got %s", l.Current())
	}
}

func TestNextPrev(t *testing.T) {
	l := newDocLinks([]string{"l1", "l2", "l3"})

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
	l := newDocLinks([]string{"l1", "l2", "l3"})

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
	l := newDocLinks([]string{"l1", "l2", "l3"})

	sc := l.GetShortcut(0)
	if sc != "A" {
		t.Errorf("expected A, got %s", l.GetShortcut(0))
	}

	sc = l.GetShortcut(3)
	if sc != "" {
		t.Errorf("expected empty string, got %s", sc)
	}

	sc = l.GetShortcut(300)
	if sc != "" {
		t.Errorf("expected empty string, got %s", sc)
	}

	sc = l.GetShortcut(300)
	if sc != "" {
		t.Errorf("expected empty string, got %s", sc)
	}
}

func TestLookupShortcuts(t *testing.T) {
	l := newDocLinks([]string{"l1", "l2", "l3"})

	sc := l.GetShortcut(1)
	if l.GetLink(sc) != "l2" {
		t.Errorf("expected l2, got %s", l.GetLink(sc))
	}

	if l.GetLink("D") != "" {
		t.Errorf("expected empty string, got %s", l.GetLink("D"))
	}
}

func TestIsActive(t *testing.T) {
	l := newDocLinks([]string{"l1", "l2", "l3"})

	l.Next()
	l.Next()

	if l.IsActive(1) != true {
		t.Errorf("expected true, got %v", l.IsActive(1))
	}
}
