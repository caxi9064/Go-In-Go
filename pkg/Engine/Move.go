package Engine

type Point uint8 	// Where pieces are places on the Go board.
type Color uint8 	// Represents the color of the piece, 0 is white 1 is black.

const (
	white Color = iota
	black
)

type Move struct {
	color Color 	// Color of piece
	xpos Point		// Move Structure, a Point is a place where you 
	ypos Point		// place a piece on the Go board.
}

func (m *Move) GetYPos() Point {
	return m.ypos
}

func (m *Move) GetXPos() Point {
	return m.xpos
}