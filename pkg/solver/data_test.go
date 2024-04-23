package solver_test

import "github.com/alendavid/puzzling/pkg/shape"

type scenario struct {
	Input    shape.Shape
	Expected shape.Moves
	Name     string
}

var (
	cases2_2 []scenario = []scenario{{
		shape.Shape{{1, 2}, {0, 3}},
		shape.Moves{{{1, 2}, {3, 0}}},
		"Move 0 once to right",
	}}
)
