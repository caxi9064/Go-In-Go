package Engine


type Board struct {
    dimension int		// Dimension of board, i.e. 19x19
	pieces [][]uint8	// Location of pieces
    moves [][]*Move     // Played moves so far.
}

func NewBoard(dim int) *Board {
	return &Board{
		dimension : dim,
		pieces: Make2D[uint8](dim, dim),
        moves: Make2D[*Move](dim,dim),
	}
}

func (b *Board) GetPieces() [][]uint8 {
    return b.pieces
}

func (b *Board) GetMoves() [][]*Move {
    return b.moves
}