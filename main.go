package main

import (
	"github.com/johnowagon/Go-in-Go/pkg/Engine"
	//"fmt"
	
)

func main() {
	e := Engine.NewEngine()
	e.GameLoop()
}