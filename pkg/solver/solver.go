package solver

import (
	"github.com/alendavid/puzzling/pkg/shape"
)

var places map[string]struct{} = make(map[string]struct{})

func Solve(input shape.Shape) shape.Moves {
	if _, ok := places[input.String()]; ok {
		return nil
	}

	places[input.String()] = struct{}{}

	columns, rows := len(input), len(input[0])

	solution := make(shape.Moves, 0)
	final := shape.Final(columns, rows)

	for _, move := range input.Next() {
		if _, ok := places[move.String()]; ok {
			continue
		}

		if move.Equal(final) {
			return append(solution, move)
		}

		steps := Solve(move)

		if len(steps) != 0 {
			solution = append(solution, move)
			return append(solution, steps...)
		}
	}

	return solution
}
