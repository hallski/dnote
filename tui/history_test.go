package tui

import (
	"reflect"
	"testing"
)

func TestCanGoBack(t *testing.T) {
	h := history[string]{}

	if h.CanGoBack() {
		t.Errorf("can go back is true, expected false")
	}

	h = history[string]{[]string{"a", "b"}, 1}

	if !h.CanGoBack() {
		t.Errorf("can go back is false, expected true")
	}

	h = history[string]{[]string{"a", "b"}, 0}
	if h.CanGoBack() {
		t.Errorf("can go back is true, expected false")
	}
}

func TestCanGoForward(t *testing.T) {
	h := history[string]{}

	if h.CanGoForward() == true {
		t.Errorf("can go forward is true, expected false")
	}

	h = history[string]{[]string{"a", "b"}, 1}

	if h.CanGoForward() == true {
		t.Errorf("can go forward is true, expected false")
	}

	h = history[string]{[]string{"a", "b"}, 0}
	if h.CanGoForward() == false {
		t.Errorf("can go forward is false, expected true")
	}
}

func TestGoBack(t *testing.T) {
	h := history[string]{[]string{"a", "b"}, 0}
	h.GoBack()
	exp := "a"
	if h.GetCurrent() != exp {
		t.Errorf("Expected curPos %s, got %s", h.GetCurrent(), exp)
	}

	h = history[string]{[]string{"a", "b"}, 1}
	h.GoBack()

	exp = "a"
	if h.GetCurrent() != exp {
		t.Errorf("Expected curPos %s, got %s", h.GetCurrent(), exp)
	}
}

func TestGoForward(t *testing.T) {
	h := history[string]{[]string{"a", "b"}, 1}
	h.GoForward()
	exp := "b"
	if h.GetCurrent() != exp {
		t.Errorf("Expected curPos %s, got %s", h.GetCurrent(), exp)
	}

	h = history[string]{[]string{"a", "b"}, 0}
	h.GoForward()

	exp = "b"
	if h.GetCurrent() != exp {
		t.Errorf("Expected curPos %s, got %s", h.GetCurrent(), exp)
	}
}

func TestPush(t *testing.T) {
	h := *NewHistory[string]()
	h.Push("a")
	h.Push("b")
	h.Push("c")
	exp := []string{"a", "b", "c"}
	if !reflect.DeepEqual(h.stack, exp) {
		t.Errorf("Expected stack %v, got %v", exp, h.stack)
	}

	h = history[string]{[]string{"a", "b"}, 2}
	h.Push("d")
	exp = []string{"a", "b", "d"}
	if !reflect.DeepEqual(h.stack, exp) {
		t.Errorf("Expected stack %v, got %v", exp, h.stack)
	}

	h = history[string]{[]string{"a", "b", "c"}, 0}
	h.Push("e")
	exp = []string{"a", "e"}
	if !reflect.DeepEqual(h.stack, exp) {
		t.Errorf("Expected stack %v, got %v", exp, h.stack)
	}
}
