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

func Final(columns, rows int) Shape {
	c := 1
	s := make(Shape, columns)

	for i := 0; i < columns; i += 1 {
		s[i] = make(Line, rows)

		for j := 0; j < rows; j += 1 {
			s[i][j] = c
			c += 1
		}
	}

	s[columns-1][rows-1] = 0

	return s
}
