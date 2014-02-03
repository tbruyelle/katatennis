package tennis

import "testing"

func TestNewGameReturnsEgalite(t *testing.T) {
	game := NewGame()

	s:=game.Score()
	if s != "0-0" {
		t.Errorf("Expected score 0-0 but was %s\n", s)
	}
}

func TestPlayerOneScores(t *testing.T) {
	game := NewGame()

	game.PlayerOneScores()
	
	s:=game.Score()
	if s != "1-0" {
		t.Errorf("Expected score 1-0 but was %s\n", s)
	}
}
