package sprint

import (
	"strconv"
	"strings"
)

func EightQueensSolver() string {
	var solutions []string
	columns := make([]int, 8)
	solve(0, columns, &solutions)
	return strings.Join(solutions, "\n")
}

func solve(row int, columns []int, solutions *[]string) {
	n := len(columns)
	if row == n {
		*solutions = append(*solutions, formatSolution(columns))
		return
	}
	for col := 0; col < n; col++ {
		if isSafe(row, col, columns) {
			columns[row] = col
			solve(row+1, columns, solutions)
		}
	}
}

func isSafe(row, col int, columns []int) bool {
	for i := 0; i < row; i++ {
		if columns[i] == col || columns[i]-col == i-row || col-columns[i] == i-row {
			return false
		}
	}
	return true
}

func formatSolution(columns []int) string {
	var solution strings.Builder
	n := len(columns)
	for i, col := range columns {
		solution.WriteString(strconv.Itoa(col + 1))
		if i < n-1 {
			solution.WriteString("")
		}
	}
	return solution.String()
}
