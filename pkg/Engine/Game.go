package Engine

import (
	//"github.com/gin-gonic/gin"

)

type Game struct {     // Main board, where pieces will live.
	board *Board
	blackScore int
	whiteScore int
}

type Point uint8 // Where pieces are places on the Go board.

func CreateGame(dim int) *Game {
	return &Game{
		board: NewBoard(dim),
	}
}

func (g *Game) PerformMove(m *Move) {
	return 
}