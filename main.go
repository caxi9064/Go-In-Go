package main

import (
	//"github.com/johnowagon/Go-in-Go/pkg/GUI"
	"github.com/johnowagon/Go-in-Go/pkg/Engine"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	

)

func main(){
	//GUI.CreateGrid();
	b := Engine.NewBoard(19)
	fmt.Println(b.GetPieces())
	a := app.NewWithID("xyz.andy.chess")
	win := a.NewWindow("Chess")

	win.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("New Game", func() {
				fmt.Println("Hello")
			})),
	))

	win.ShowAndRun()
}