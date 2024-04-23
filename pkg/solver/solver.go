package solver

import (
	"github.com/alendavid/puzzling/pkg/shape"
)

func Solve(input shape.Shape) shape.Moves {
	columns, rows := len(input), len(input[0])

	solution := make(shape.Moves, 0)
	final := shape.Final(columns, rows)

	if input.Equal(final) == 0 {
		return solution
	}

	solution = append(solution, input.Steps(final)...)

	return solution

}
