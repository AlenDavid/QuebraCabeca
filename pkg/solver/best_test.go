package solver_test

import (
	"testing"

	"github.com/alendavid/puzzling/pkg/shape"
	"github.com/alendavid/puzzling/pkg/solver"
	"github.com/stretchr/testify/assert"
)

func TestBestSearch(t *testing.T) {
	input := shape.Shape{
		{2, 1, 3},
		{8, 0, 4},
		{7, 5, 6},
	}

	res := solver.BestSearch(input)

	assert.Equal(t, shape.Final(), res[len(res)-1])
}

func BenchmarkBestSearch(b *testing.B) {
	input := shape.Shape{
		{2, 1, 3},
		{8, 0, 4},
		{7, 5, 6},
	}

	for i := 0; i < b.N; i++ {
		res := solver.BestSearch(input)

		assert.Equal(b, shape.Final(), res[len(res)-1])
	}
}
