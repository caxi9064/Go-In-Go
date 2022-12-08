package Engine

import (
   //"fmt"
)

type piece interface { //component
	getPieceStats() string
}

// white pieces 
type whitePiece struct { //concrete component
}
func (p *whitePiece) getPieceStats() string {
	return "White: "
}
//black pieces
type blackPiece struct { //concrete component
}
func (p *blackPiece) getPieceStats() string { // getPieceColor
	return "Black: "
}

type displayBoard struct { //concrete decorator
	piece piece
}
func (c *displayBoard) getPieceStats() string {
	pieceColor := c.piece.getPieceStats()
	return pieceColor + "7"
}
type gameStats struct {
	piece piece
}
func (c *gameStats) getPieceStats() string { 
	pieceColor := c.piece.getPieceStats()
	return pieceColor + "10"
}