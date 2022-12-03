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
	for {
		e.game.board.DrawBoard()

		fmt.Println("\nEnter your move coords separated by a space:")
		var xpos Point = 0
		var ypos Point = 0
		_, errorf := fmt.Scanln(&xpos, &ypos)
		if errorf != nil {
			fmt.Printf("Input could not be read: %v\n", errorf)
			continue
		}
		if counter % 2 == 1 {
			c = black
		}else{
			c = white
		}
		move, err := getMove(xpos, ypos, c)
		if errorf != nil {
			fmt.Printf("Move could not be played: %v\n", err)
			continue
		}
		e.game.PerformMove(move)
		counter++
		fmt.Printf("%s played the move %d, %d.\n", EnumString(c), xpos, ypos)
		if counter == 5 {
			e.game.board.Capture(c)
		}
	}
}