package tui

type history[T interface{}] struct {
	stack  []T
	curPos int
}

func NewHistory[T interface{}]() *history[T] {
	return &history[T]{[]T{}, -1}
}

func (h *history[T]) CanGoBack() bool {
	if h.curPos > 0 && len(h.stack) > h.curPos {
		return true
	}
	return false
}

func (h *history[T]) CanGoForward() bool {
	if h.curPos < len(h.stack)-1 && h.curPos >= 0 {
		return true
	}
	return false
}

func (h *history[T]) GoBack() T {
	if h.CanGoBack() {
		h.curPos--
	}

	return h.GetCurrent()
}

func (h *history[T]) GoForward() T {
	if h.CanGoForward() {
		h.curPos++
	}

	return h.GetCurrent()
}

func (h *history[T]) GetCurrent() T {
	return h.stack[h.curPos]
}

func (h *history[T]) Push(t T) {
	if len(h.stack) == 0 {
		h.stack = append(h.stack, t)
		h.curPos = 0
		return
	}

	if h.curPos < len(h.stack)-1 {
		h.stack = append(h.stack[0:h.curPos+1], t)
	} else {
		h.stack = append(h.stack, t)
	}
	h.curPos++
}
