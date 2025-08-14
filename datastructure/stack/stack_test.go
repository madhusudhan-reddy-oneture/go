package stack_test

import (
	"testing"

	"github.com/madhusudhan-reddy-oneture/gotbd/datastructure/stack"
)

func TestStack(t *testing.T) {
	myStack := stack.NewStack[int]()

	// check stack is empty
	AssertTrue(t, myStack.IsEmpty())

	// add a thing, then check it's not empty
	myStack.Push(1234)
	AssertFalse(t, myStack.IsEmpty())

	// check if the top returns the correct value
	value, ok := myStack.Top()
	AssertTrue(t, ok)
	AssertEqual(t, value, 1234)

	// pop from the stack
	value, ok = myStack.Pop()
	AssertTrue(t, ok)
	AssertEqual(t, value, 1234)

	// check if the satck is empty
	AssertTrue(t, myStack.IsEmpty())

	// top command should return false beacuse the stack is empty
	_, ok = myStack.Top()
	AssertFalse(t, ok)
}

func AssertTrue(t testing.TB, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t testing.TB, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}
