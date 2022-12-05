package Engine

/*
 *	Factory pattern to handle the creation of moves.
 *
 */

import (
	//"fmt"
)

func getMove(x, y Point, c Color) (*Move, error) {
	return &Move{
		color: c,
		xpos: x,
		ypos: y,
	}, nil
}