package Engine

import (
	"fmt"

)

type Game struct {     
	board *Board
	blackScore int
	whiteScore int
}



func CreateGame() *Game {
	fmt.Println("Welcome to Go in Go!")
	fmt.Println("Input number for board size (X by X): ")
	var input int
	fmt.Scanln(&input)

	return &Game{
		board: NewBoard(input),
	}
}

