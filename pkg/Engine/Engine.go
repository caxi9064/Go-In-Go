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

	//STRATEGY PATTERN
	add := Operation{Addition{}}
	var i int = add.Operate(3, 5) // 8
	fmt.Println(i)

	mult := Operation{Multiplication{}}
	var j int = mult.Operate(3, 5) // 15
	fmt.Println(j)

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