package solver

import (
	"github.com/alendavid/puzzling/pkg/moves"
	"github.com/alendavid/puzzling/pkg/shape"
)

func BestSearch(input shape.Shape) moves.Moves {
	final := shape.Final()

	open := make(moves.Moves, 0)
	close := make(moves.Moves, 0)

	open = append(open, input)

	for len(open) > 0 {
		x := open[0]
		open = open[1:]

		if x.Equal(final) {
			return append(make(moves.Moves, 0), input, x)
		}

		for _, child := range x.Children() {
			if !open.Contains(child) && !close.Contains(child) {
				open = append(open, child)
			}
		}

		close = append(close, x)

		open.Sort()
	}

	return nil
}
