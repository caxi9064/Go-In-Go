package Engine

//import "fmt"


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
		if counter % 2 == 1 {
			c = black
		}else{
			c = white
		}
		e.game.board.PromptMove(c)
		counter++
	}
}