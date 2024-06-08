package solver

import (
	"github.com/alendavid/puzzling/pkg/shape"
)

func WideSearch(input shape.Shape) shape.Moves {
	final := shape.Final()

	open := make(shape.Moves, 0)
	close := make(shape.Moves, 0)

	open = append(open, input)

	for len(open) > 0 {
		x := open[0]
		open = open[1:]

		if x.Equal(final) {
			return append(make(shape.Moves, 0), input, x)
		}

		close = append(close, x)

		for _, move := range x.Next() {
			if !open.Contains(move) && !close.Contains(move) {
				open = append(open, move)
			}
		}
	}

	return nil
}
