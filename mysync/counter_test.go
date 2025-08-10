package mysync

import (
	"math/rand"
	"sync"
	"testing"
)

func getRandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func TestCounter(t *testing.T) {
	t.Run("incrementing counter n times leave it at n", func(t *testing.T) {
		n := getRandomInt(5, 10)

		counter := NewCounter()
		for range n {
			counter.Inc()
		}

		assertCounter(t, counter, n)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for range wantedCount {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
