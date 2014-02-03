package tennis

import (
	"strconv"
)

type Game struct {
	playerOneScore int
}

func (g *Game) Score() string {
	return strconv.FormatInt(int64(g.playerOneScore), 10) + "-0"
}

func (g *Game) PlayerOneScores() {
	g.playerOneScore++
}

func NewGame() *Game {
	return &Game{}
}
