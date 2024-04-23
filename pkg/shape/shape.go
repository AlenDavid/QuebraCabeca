package shape

type Line []int
type Shape []Line
type Moves []Shape

func (s Shape) zero() (int, int) {
	for i := 0; i < len(s); i += 1 {
		for j := 0; j < len(s[0]); j += 1 {
			if s[i][j] == 0 {
				return i, j
			}
		}
	}

	return -1, -1
}
