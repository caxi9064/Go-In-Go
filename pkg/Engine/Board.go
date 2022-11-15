package Engine

type Board struct {     // Main board, where pieces will live.
	dimension int		// Dimension of board, i.e. 19x19
	pieces [][]uint8	// Location of pieces
    moves [][]*Move     // Played moves so far.
}

type Point uint8 // Where pieces are places on the Go board.

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

func Make2D[T any](n, m int) [][]T {
    /* Returns a 2d slice of a given type. */
    matrix := make([][]T, n)
    rows := make([]T, n*m)
    for i, startRow := 0, 0; i < n; i, startRow = i+1, startRow+m {
        endRow := startRow + m
        matrix[i] = rows[startRow:endRow:endRow]
    }
    return matrix
}