package Engine

import (
	"fmt"
)

type Board struct {
    dimension int			// Dimension of board, i.e. 19x19
	pieces [][]Color		// Location of pieces
	pastPosition [][]Color 	// To avoid the Ko rule.
    moves []*Move     		// Played moves so far.
}

func NewBoard(dim int) *Board {
	return &Board{
		dimension : dim,
		pieces: Make2D[Color](dim, dim),
        moves: make([]*Move, dim), // A slice of moves
	}
}

func (b *Board) GetPieces() [][]Color {
    return b.pieces
}

func (b *Board) GetMoves() []*Move {
    return b.moves
}

func (b *Board) DrawBoard() {
	// Draw the board to the console. Will be called after each turn.
	for i := 0; i < b.dimension; i++ {
		if (i == 0 ){
			for k := 0; k < b.dimension; k++ {
				if k < 10 {
					if k == 0 {
						fmt.Print("   ")
					}
					fmt.Print(" " + fmt.Sprint(k) + " ")
				}else{
					fmt.Print(fmt.Sprint(k) + " ")
				}
			}
		}
		fmt.Println()
		if i+1 < 10 {
			fmt.Print(" " + fmt.Sprint(i) + " ")
		}else{
			fmt.Print(fmt.Sprint(i) + " ")
		}
		for j := 0; j < b.dimension; j++ {
			if b.pieces[i][j] == white {
				fmt.Print(" ○ ")
			}else if b.pieces[i][j] == black{
				fmt.Print(" ● ")
			}else{
				fmt.Print(" + ")
			}
		}
	}
}

func (b *Board) IsLegalMove(m *Move) bool {
	// Checks to see if the move is legal.
	if b.pieces
	return false
}

func (b *Board) Capture(c Color){
	// Removes captured pieces of the opposite color from the board
	oppColor := GetOppositeColor(c)
	for i,row := range b.pieces {
		for j,val := range row {
			if val == oppColor {
				fmt.Printf("Black moved: %d, %d", i, j)
			}
		}
	}
}