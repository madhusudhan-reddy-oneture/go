package math

import (
	"slices"
	"testing"
)

func TestSumAllTails(t *testing.T) {

	checkSum := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Sum of all tails", func(t *testing.T) {
		got := SumAllTails([]int{2, 3}, []int{0, 9})
		want := []int{3, 9}
		checkSum(t, got, want)
	})

	t.Run("safely run empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSum(t, got, want)
	})
}

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !slices.Equal(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}
