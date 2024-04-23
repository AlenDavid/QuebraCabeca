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

	t.Run("can find zero", func(t *testing.T) {
		cases := []scenario{
			{"1,1", Shape{{1, 2}, {3, 0}}, [2]int{1, 1}},
			{"0,1", Shape{{1, 2}, {0, 3}}, [2]int{0, 1}},
		}

		for _, c := range cases {
			t.Run(c.Name, func(t *testing.T) {
				i, j := c.Input.zero()

				assert.Equal(t, c.Expected, [2]int{i, j})
			})
		}
	})

	t.Run("cannot find zero", func(t *testing.T) {
		cases := []scenario{
			{"", Shape{{1, 2}, {3, 4}}, [2]int{-1, -1}},
		}

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

func TestMoves(t *testing.T) {
	t.Run("moving up", func(t *testing.T) {
		input := Shape{
			{1, 2},
			{3, 0},
		}

		expected := Shape{
			{1, 0},
			{3, 2},
		}

		worked, up := input.Up()

		assert.True(t, worked)
		assert.Equal(t, expected, up)
		assert.NotEqual(t, input, up)
	})

	t.Run("moving down", func(t *testing.T) {
		input := Shape{
			{1, 0},
			{3, 2},
		}

		expected := Shape{
			{1, 2},
			{3, 0},
		}

		worked, up := input.Down()

		assert.True(t, worked)
		assert.Equal(t, expected, up)
		assert.NotEqual(t, input, up)
	})

	t.Run("moving left", func(t *testing.T) {
		input := Shape{
			{1, 0},
			{3, 2},
		}

		expected := Shape{
			{0, 1},
			{3, 2},
		}

		worked, up := input.Left()

		assert.True(t, worked)
		assert.Equal(t, expected, up)
		assert.NotEqual(t, input, up)
	})
}
