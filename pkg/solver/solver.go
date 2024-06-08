package solver

import (
	"github.com/alendavid/puzzling/pkg/shape"
)

func Solve(input shape.Shape) shape.Moves {
	return BestSearch(input)
}
