package poker_test

import (
	"strings"
	"testing"

	"github.com/madhusudhan-reddy-oneture/gotbd/my-app/poker"
)

func TestCLI(t *testing.T) {

	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")

		playerStore := poker.NewStubPlayerStore()

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")

		playerStore := poker.NewStubPlayerStore()

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

}
