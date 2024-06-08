package solver

import (
	"slices"

	"github.com/alendavid/puzzling/pkg/moves"
	"github.com/alendavid/puzzling/pkg/shape"
	"github.com/alendavid/puzzling/pkg/state"
)

func StateContains(st []state.State, value shape.Shape) bool {
	return slices.ContainsFunc(st, func(s state.State) bool {
		return s.Equal(value)
	})
}

func StateSort(st []state.State) {
	final := shape.Final()

	slices.SortFunc(st, func(a, b state.State) int {
		return a.Distance(final) - b.Distance(final)
	})
}

func BestSearch(input shape.Shape) moves.Moves {
	final := shape.Final()

	open := make([]state.State, 0)
	close := make([]state.State, 0)

	open = append(open, state.Begin(input))

	for len(open) > 0 {
		x := open[0]
		open = open[1:]

		if x.Equal(final) {
			return x.Moves()
		}

		for _, child := range x.Shape.Children() {
			if !StateContains(open, child) && !StateContains(close, child) {
				open = append(open, x.Next(child))
			}
		}

		close = append(close, x)

		StateSort(open)
	}

	return nil
}
