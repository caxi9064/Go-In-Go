package Engine

import (
	"fmt"
)

/*
 *	Command pattern implementation, used for handling user move options.
 */

type Command interface {
	// Base interface
	execute()
}

type Invoker struct {
	// Invoker for each cmd
	command Command
}

// Players can either move a piece or pass their turn.

type MoveCommand struct {
	move *Move
	board *Board
}

type PassCommand struct {
	c Color
}

func (M *MoveCommand) execute() {
	M.board.PerformMove(M.move)
}

func (P *PassCommand) execute() {
	fmt.Println(EnumString(P.c) + " has passed their turn.")
}