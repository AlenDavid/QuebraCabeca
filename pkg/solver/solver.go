package solver

type shape interface {
	~[][]int
}

func Done(a, b int) [][]int {
	c := 1
	s := make([][]int, a)

	for i := 0; i < a; i += 1 {
		s[i] = make([]int, b)

		for j := 0; j < b; j += 1 {
			s[i][j] = c
			c += 1
		}
	}

	s[a-1][b-1] = 0

	return s
}

func Step[S shape](a, b S) S {
	return a
}

func Solve[S shape](input S) []S {
	a, b := len(input), len(input[0])

	solution := make([]S, 0)

	solution = append(solution, Done(a, b))

	return solution
}
