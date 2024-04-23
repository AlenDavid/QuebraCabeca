package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZero(t *testing.T) {
	type scenario struct {
		Name     string
		Input    Shape
		Expected [2]int
	}

	cases := []scenario{
		{"1,1", Shape{{1, 2}, {3, 0}}, [2]int{1, 1}},
		{"1,0", Shape{{1, 2}, {0, 3}}, [2]int{1, 0}},
	}

	t.Run("can find zero", func(t *testing.T) {
		for _, c := range cases {
			t.Run(c.Name, func(t *testing.T) {
				i, j := c.Input.zero()

				assert.Equal(t, c.Expected, [2]int{i, j})
			})
		}
	})
}

func TestFinal(t *testing.T) {
	t.Run("returns the correct shape for any given XY puzzle", func(t *testing.T) {
		t.Run("2x2 puzzle", func(t *testing.T) {
			expected := Shape{{1, 2}, {3, 0}}
			assert.Equal(t, expected, Final(2, 2))
		})

		t.Run("3x3 puzzle", func(t *testing.T) {
			expected := Shape{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}
			assert.Equal(t, expected, Final(3, 3))

		})
	})
}
