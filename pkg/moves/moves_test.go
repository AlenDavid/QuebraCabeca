package moves

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoves(t *testing.T) {
	t.Run("sorting", func(t *testing.T) {
		input := Moves{
			{{2, 1, 3},
				{8, 6, 4},
				{7, 0, 5}},
			{{2, 1, 3},
				{8, 0, 4},
				{7, 6, 5}},
			{{1, 2, 3},
				{8, 0, 4},
				{7, 6, 5}},
		}

		expected := Moves{
			{{1, 2, 3},
				{8, 0, 4},
				{7, 6, 5}},
			{{2, 1, 3},
				{8, 0, 4},
				{7, 6, 5}},
			{{2, 1, 3},
				{8, 6, 4},
				{7, 0, 5}},
		}

		input.Sort()
		assert.Equal(t, input, expected)
	})
}
