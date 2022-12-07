package Engine

import "fmt"


type Engine struct {
	game *Game
}

func NewEngine() *Engine {
	g := CreateGame()
	return &Engine{
		game: g,
	}
}

func (e *Engine) GameLoop() {
	var counter int = 0
	var c Color

	//DECORATOR PATTERN
	blackPiece := &blackPiece{}

	//Add cheese topping
	blackPieceWithGameStats:= &gameStats{
		piece: blackPiece,
	}

	//Add tomato topping
	blackPieceWithGameStatsAndDisplay := &displayBoard{
		piece: blackPieceWithGameStats,
	}

	fmt.Printf(blackPieceWithGameStatsAndDisplay.getPieceStats())
	fmt.Printf("\n")

	whitePiece := &whitePiece{}

	//Add cheese topping
	whitePieceWithGameStats := &gameStats{
		piece: whitePiece,
	}

	fmt.Printf(whitePieceWithGameStats.getPieceStats())
	fmt.Printf("\n")
	//edit naming

	for {
		if counter % 2 == 1 {
			c = black
		}else{
			c = white
		}
		e.game.board.PromptMove(c)
		counter++
	}
}