package tennis

import "testing"

var game *Game

func setup() {
	game = NewGame("Agassi", "Djoko")
}

func assertScore(t *testing.T, expected string) {
	s := game.Score()
	if s != expected {
		t.Errorf("Expected score \"%s\" but was \"%s\"", expected, s)
	}
}

func createScore(scorePlayerOne, scorePlayerTwo int) {
	for i := 0; i < scorePlayerOne; i++ {
		game.PlayerOneScores()
	}
	for i := 0; i < scorePlayerTwo; i++ {
		game.PlayerTwoScores()
	}
}
func TestNewGameReturnsEgalite(t *testing.T) {
	setup()

	assertScore(t, "0-0")
}

func TestPlayerOneScores(t *testing.T) {
	setup()

	createScore(1, 0)

	assertScore(t, "15-0")
}

func TestPlayerOneScores2(t *testing.T) {
	setup()

	createScore(2, 0)

	assertScore(t, "30-0")
}

func TestPlayerOneScores3(t *testing.T) {
	setup()

	createScore(3, 0)

	assertScore(t, "40-0")
}
func TestPlayerOneWins(t *testing.T) {
	setup()

	createScore(4, 0)

	assertScore(t, "Agassi wins")
}

func TestPlayerTwoScores(t *testing.T) {
	setup()

	createScore(0, 1)

	assertScore(t, "0-15")
}

func TestPlayerTwoScores2(t *testing.T) {
	setup()

	createScore(0, 2)

	assertScore(t, "0-30")
}

func TestPlayerTwoScores3(t *testing.T) {
	setup()

	createScore(0, 3)

	assertScore(t, "0-40")
}

func TestPlayerTwoWins(t *testing.T) {
	setup()

	createScore(0, 4)

	assertScore(t, "Djoko wins")
}

func TestFifteenAll(t *testing.T) {
	setup()

	createScore(1, 1)

	assertScore(t, "15-15")
}

func TestDeuce(t *testing.T) {
	setup()

	createScore(3, 3)

	assertScore(t, "Deuce")
}

func TestPlayerOneAdvantage(t *testing.T) {
	setup()

	createScore(4, 3)

	assertScore(t, "Advantage Agassi")
}

func TestPlayerTwoAdvantage(t *testing.T) {
	setup()

	createScore(3, 4)

	assertScore(t, "Advantage Djoko")
}

func TestDeuceAfterAdvantage(t *testing.T) {
	setup()

	createScore(4, 4)

	assertScore(t, "Deuce")
}

func TestPlayerOneWinsAfterAdvantage(t *testing.T) {
	setup()

	createScore(6, 4)

	assertScore(t, "Agassi wins")
}

func TestPlayerTwoWinsAfterAdvantage(t *testing.T) {
	setup()

	createScore(7, 9)

	assertScore(t, "Djoko wins")
}
