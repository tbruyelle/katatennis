package tennis

import "fmt"

type Game struct {
	scorePlayerOne int
	namePlayerOne  string
	scorePlayerTwo int
	namePlayerTwo  string
}

func (g *Game) Score() string {
	if g.isDeuce() {
		return "Deuce"
	}
	if ok, player := g.hasAdvantage(); ok {
		return "Advantage " + player
	}
	if ok, player := g.hasWinner(); ok {
		return player + " wins"
	}

	return translateScore(g.scorePlayerOne) + "-" + translateScore(g.scorePlayerTwo)
}

func (g *Game) isDeuce() bool {
	return g.scorePlayerOne == g.scorePlayerTwo && g.scorePlayerOne >= 3
}

func (g *Game) hasAdvantage() (bool, string) {
	if g.scorePlayerOne > 3 && g.scorePlayerOne == g.scorePlayerTwo+1 {
		return true, g.namePlayerOne
	}
	if g.scorePlayerTwo > 3 && g.scorePlayerTwo == g.scorePlayerOne+1 {
		return true, g.namePlayerTwo
	}
	return false, ""
}

func (g *Game) hasWinner() (bool, string) {
	if g.scorePlayerOne > 3 && g.scorePlayerOne >= g.scorePlayerTwo+2 {
		return true, g.namePlayerOne
	}
	if g.scorePlayerTwo > 3 && g.scorePlayerTwo >= g.scorePlayerOne+2 {
		return true, g.namePlayerTwo
	}
	return false, ""
}

func translateScore(score int) string {
	switch score {
	case 0:
		return "0"
	case 1:
		return "15"
	case 2:
		return "30"
	case 3:
		return "40"
	}
	return fmt.Sprintf("## Unhandled score %d ##", score)
}

func (g *Game) PlayerOneScores() {
	g.scorePlayerOne++
}

func (g *Game) PlayerTwoScores() {
	g.scorePlayerTwo++
}

func NewGame(playerOne, playerTwo string) *Game {
	return &Game{namePlayerOne: playerOne, namePlayerTwo: playerTwo}
}
