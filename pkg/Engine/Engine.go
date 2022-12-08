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

	blackPieceWithGameStats:= &gameStats{
		piece: blackPiece,
	}

	fmt.Printf("Player Stats\n")
	fmt.Printf(blackPieceWithGameStats.getPieceStats())
	fmt.Printf("\n")

	whitePiece := &whitePiece{}
	whitePieceWithGameStats := &gameStats{
		piece: whitePiece,
	}

	fmt.Printf(whitePieceWithGameStats.getPieceStats())
	fmt.Printf("\n")

	//STRATEGY PATTERN
	add := Operation{AddPieces{}}
	var i int = add.Operate(1, 3) // 8
	fmt.Println("Pieces Captured: ", i)

	subt := Operation{SubtractPieces{}}
	var j int = subt.Operate(10, 3) // 15
	fmt.Println("Pieces Left: ", j)
	fmt.Printf("\n")
	fmt.Printf("--------------------------------\n")


	//change for loop to while 
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