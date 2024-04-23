package solver

import (
	"slices"

	"github.com/alendavid/puzzling/pkg/shape"
)

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
	final := shape.Final(columns, rows)

	if slices.CompareFunc(input, final, lineCompare) == 0 {
		return solution
	}

	solution = append(solution, Steps(input, final)...)

	return solution

}
