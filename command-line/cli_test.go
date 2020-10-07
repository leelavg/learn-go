package poker_test

import (
	"strings"
	"testing"

	poker "github.com/leelavg/learn-go/command-line"
)

func TestCLI(t *testing.T) {

	t.Run("record leela win from user input", func(t *testing.T) {
		in := strings.NewReader("Leela wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Leela")

	})

	t.Run("record Neela win from users input", func(t *testing.T) {
		in := strings.NewReader("Neela wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Neela")
	})
}
