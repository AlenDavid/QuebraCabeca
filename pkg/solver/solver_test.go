package solver_test

import (
	"testing"

	"github.com/alendavid/puzzling/pkg/solver"
	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	t.Run("solve 2x2 cases", func(t *testing.T) {
		for _, puzzle := range cases2_2 {
			t.Run(puzzle.Name, func(t *testing.T) {
				result := solver.Solve(puzzle.Input)

				assert.Equal(t, puzzle.Expected, result)
			})
		}
	})
}
