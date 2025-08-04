package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(34, 35)
	expected := 69

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(400, 20)
	fmt.Println(sum)
	// Output: 420
}
