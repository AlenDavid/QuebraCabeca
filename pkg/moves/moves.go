package moves

import (
	"fmt"
	"slices"

	"github.com/alendavid/puzzling/pkg/shape"
)

type Moves []shape.Shape

func (m Moves) String() string {
	t := ""

	for _, shape := range m {
		t += fmt.Sprintf("%s\n", shape.String())
	}

	return t
}

func (m Moves) Contains(a shape.Shape) bool {
	return slices.ContainsFunc(m, func(b shape.Shape) bool {
		return a.Equal(b)
	})
}

func (m Moves) Sort() {
	final := shape.Final()

	slices.SortFunc(m, func(a, b shape.Shape) int {
		return a.Distance(final) - b.Distance(final)
	})
}
