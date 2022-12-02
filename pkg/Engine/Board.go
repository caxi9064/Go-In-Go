package Engine

import (
	"fmt"
)

type Board struct {
    dimension int		// Dimension of board, i.e. 19x19
	pieces [][]Color	// Location of pieces
    moves [][]*Move     // Played moves so far.
}

func NewBoard(dim int) *Board {
	return &Board{
		dimension : dim,
		pieces: Make2D[Color](dim, dim),
        moves: Make2D[*Move](dim,dim),
	}
}

func (b *Board) GetPieces() [][]Color {
    return b.pieces
}

func (b *Board) GetMoves() [][]*Move {
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
				fmt.Print(" o ")
			}else if b.pieces[i][j] == black{
				fmt.Print(" * ")
			}else{
				fmt.Print(" + ")
			}
		}
	}
}