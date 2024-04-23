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

	for _, move := range input.Next() {
		if move.Equal(final) == 0 {
			return append(solution, move)
		}
	}

	return solution
}
