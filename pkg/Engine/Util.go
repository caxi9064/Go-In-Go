package Engine


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