package solver

import (
	"github.com/alendavid/puzzling/pkg/shape"
)

var places map[string]struct{} = make(map[string]struct{})

func wideSearch(input shape.Shape) shape.Moves {
	if _, ok := places[input.String()]; ok {
		return nil
	}

	places[input.String()] = struct{}{}

	solution := make(shape.Moves, 0)
	final := shape.Final()

	for _, move := range input.Next() {
		if _, ok := places[move.String()]; ok {
			continue
		}

		if move.Equal(final) {
			return append(solution, move)
		}

		steps := wideSearch(move)

		if len(steps) != 0 {
			solution = append(solution, move)
			return append(solution, steps...)
		}
	}

	return solution
}
