package solver_test

type scenario struct {
	Input    [][]int
	Expected [][][]int
	Name     string
}

var (
	cases2_2 []scenario = []scenario{{
		[][]int{{1, 2}, {0, 3}},
		[][][]int{{{1, 2}, {3, 0}}},
		"Move 0 once to right",
	}}
)
