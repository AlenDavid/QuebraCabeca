package solver

import (
	"github.com/alendavid/puzzling/pkg/moves"
	"github.com/alendavid/puzzling/pkg/shape"
)

func Solve(input shape.Shape) moves.Moves {
	return BestSearch(input)
}
