package solver

import (
	"slices"

	"github.com/alendavid/puzzling/pkg/shape"
)

func Final(columns, rows int) shape.Shape {
	c := 1
	s := make(shape.Shape, columns)

	for i := 0; i < columns; i += 1 {
		s[i] = make(shape.Line, rows)

		for j := 0; j < rows; j += 1 {
			s[i][j] = c
			c += 1
		}
	}

	s[columns-1][rows-1] = 0

	return s
}

func Moves(a shape.Shape) shape.Moves {
	return nil
}

func Steps(a, b shape.Shape) shape.Moves {
	steps := make(shape.Moves, 0)

	return steps
}

func lineCompare(e1, e2 shape.Line) int {
	return slices.Compare(e1, e2)
}

func Solve(input shape.Shape) shape.Moves {
	columns, rows := len(input), len(input[0])

	solution := make(shape.Moves, 0)
	final := Final(columns, rows)

	if slices.CompareFunc(input, final, lineCompare) == 0 {
		return solution
	}

	solution = append(solution, Steps(input, final)...)

	return solution

}
