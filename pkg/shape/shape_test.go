package shape

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type scenario struct {
	Name     string
	Input    Shape
	Expected [2]int
}

func TestZero(t *testing.T) {
	cases := []scenario{
		{"1,1", Shape{{1, 2}, {3, 0}}, [2]int{1, 1}},
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
