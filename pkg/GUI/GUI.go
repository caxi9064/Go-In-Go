package GUI

import (
	"log"
	"fmt"
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	//"fyne.io/fyne/v2/position"
	"fyne.io/fyne/v2/data/binding"
	//"github.com/johnowagon/Go-in-Go/pkg/Engine"
)

type GUI struct { // How our GUI will be represented
	grid *fyne.Container
		

	blackTurn binding.Bool
}

func CreateGrid(dim int) {
	/*
	 * Will create and return the playing grid for our playing surface. 
	 * Maybe we need to define a struct for our GUI elements; grid, pieces, pass button, etc.
	 */
	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Layout")

	grid := container.NewGridWithColumns(dim)
	//This line will likely not work on your computer catherine
	image := canvas.NewImageFromFile("/Users/johnaldrete/CSCI4448/Go-in-Go/pkg/imgs/9x9board.png")

	for y := 0; y < dim; y++{
		for x := 0; x < dim; x++{
			
			
			button := widget.NewButton(fmt.Sprint(x) + "," + fmt.Sprint(y), func () {
				log.Printf("Tapped %d %d", x, y)
			})
			//button.Move(fyne.NewPos(float32(x*40), float32(y*40)))
			//bg := canvas.NewRectangle(color.Gray{0xE0})

			bg := container.NewMax(canvas.NewRectangle(color.RGBA{0,0,0,0}), button)
			//button.Move(fyne.NewPos(float32(x*40), float32(y*40)))
			grid.Add(bg)
		}
	}
	
	myWindow.SetContent(container.NewMax(grid, image))
	myWindow.Resize(fyne.NewSize(700, 700))
	myWindow.ShowAndRun()
}