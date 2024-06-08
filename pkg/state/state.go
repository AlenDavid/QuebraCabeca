package state

import (
	"slices"

	"github.com/alendavid/puzzling/pkg/moves"
	"github.com/alendavid/puzzling/pkg/shape"
)

type State struct {
	shape.Shape

	Previous *State
	Size     int
}

func Begin(shape shape.Shape) State {
	return State{shape, nil, 0}
}

func (s State) Next(shape shape.Shape) State {
	return State{shape, &s, s.Size + 1}
}

func (s *State) Moves() moves.Moves {
	arr := make(moves.Moves, 0)

	for t := s; t != nil; t = t.Previous {
		arr = append(arr, t.Shape)
	}

	slices.Reverse(arr)

	return arr
}
