package stack

type stack[t any] struct {
	values []t
}

func NewStack[T any]() *stack[T] {
	return new(stack[T])
}

func (s *stack[t]) Push(value t) {
	s.values = append(s.values, value)
}

func (s *stack[t]) Top() (t, bool) {
	if s.IsEmpty() {
		var zero t
		return zero, false
	}
	return s.values[len(s.values)-1], true
}

func (s *stack[t]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *stack[t]) Pop() (t, bool) {
	el, ok := s.Top()
	if ok {
		last := len(s.values) - 1
		s.values = s.values[:last]
	}
	return el, ok
}
