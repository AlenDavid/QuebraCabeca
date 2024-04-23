package solver

import "slices"

type line = []int
type shape = []line
type steps = []shape

func Final(columns, rows int) shape {
	c := 1
	s := make(shape, columns)

	for i := 0; i < columns; i += 1 {
		s[i] = make(line, rows)

		for j := 0; j < rows; j += 1 {
			s[i][j] = c
			c += 1
		}
	}

	s[columns-1][rows-1] = 0

	return s
}

func Moves(a shape) []shape {
	return make([][][]int, 0)
}

func Steps(a, b shape) []shape {
	steps := make([]shape, 0)

	return steps
}

func lineCompare(e1, e2 line) int {
	return slices.Compare(e1, e2)
}

func Solve(input shape) steps {
	columns, rows := len(input), len(input[0])

	solution := make(steps, 0)
	final := Final(columns, rows)

	if slices.CompareFunc(input, final, lineCompare) == 0 {
		return solution
	}

	solution = append(solution, Steps(input, final)...)

	return solution

}
