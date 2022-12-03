package Engine

import (
    "fmt"
)

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

func EnumString(c Color) string {
    // Return the string of the Color enum.
    switch c {
        case 1 : return "○White"
        case 2 : return "●Black"
        default : return fmt.Sprint(c)
    }
}

func GetOppositeColor(c Color) Color {
    switch c {
        case 1 : return black
        case 2 : return white
        default : return 0
    }
}

func Compare2d(a, b[][]*Color) bool {
    if len(a) != len(b) {
        return false
    }
    for i,v := range a {
        for j,v2 := range v {
            if b[i][j] != v2 {
                return false
            }
        }
    }
    return true
}