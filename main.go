package main

import (
	"github.com/johnowagon/Go-in-Go/pkg/Engine"
	"fmt"
	//"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/app"
	

)

func main() {
	eng := Engine.CreateGameEngine(19)
	fmt.Print(eng.game)
}