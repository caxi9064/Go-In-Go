package Engine

import (
	//"net/http"
	"github.com/gin-gonic/gin"
)

type GameEngine struct {
	game *Game
	router *gin.Engine
}

func CreateGameEngine(dim int) *GameEngine {
	r := gin.Default()
	g := CreateGame(dim)
	ge := GameEngine{game:g, router:r}
	return &ge
}



