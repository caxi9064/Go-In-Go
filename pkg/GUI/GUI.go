package GUI

import (
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/position"
	"fyne.io/fyne/v2/data/binding"
)

type GUI struct { // How our GUI will be represented
	grid *board
	

	blackTurn binding.Bool
}

func CreateGrid() {
	/*
	 * Will create and return the playing grid for our playing surface. 
	 * Maybe we need to define a struct for our GUI elements; grid, pieces, pass button, etc.
	 */
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Layout")

	grid := container.NewGridWithColumns(19)

	for y := 0; y < 19; y++{
		for x := 0; x < 19; x++{
			bg := canvas.NewRectangle(color.Gray{0xE0})
			
			grid.Add(bg)
		}
	}
	
	myWindow.SetContent(grid)
	myWindow.Resize(fyne.NewSize(700, 700))
	myWindow.ShowAndRun()
}