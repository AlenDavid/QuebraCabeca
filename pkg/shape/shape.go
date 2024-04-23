package shape

import "slices"

type Line []int
type Shape []Line
type Moves []Shape

func (s Shape) zero() (x int, y int) {
	for y = 0; y < len(s); y += 1 {
		for x = 0; x < len(s[0]); x += 1 {
			if s[y][x] == 0 {
				return x, y
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

func (a Shape) Columns() int {
	return len(a)
}

func (a Shape) Rows() int {
	return len(a[0])
}

func (a Shape) Copy() Shape {
	b := make(Shape, len(a))

	for i := 0; i < len(a[0]); i += 1 {
		b[i] = append(b[i], a[i]...)
	}

	return b
}

func (a Shape) Up() (bool, Shape) {
	if len(a) < 2 {
		return false, a
	}

	b := a.Copy()

	x, y := b.zero()

	if y == 0 {
		return false, b
	}

	v := b[y-1][x]
	b[y-1][x] = 0
	b[y][x] = v

	return true, b
}

func (a Shape) Down() (bool, Shape) {
	if len(a) < 2 {
		return false, a
	}

	b := a.Copy()

	x, y := b.zero()

	if y == len(a) {
		return false, b
	}

	v := b[y+1][x]
	b[y+1][x] = 0
	b[y][x] = v

	return true, b
}

func (a Shape) Left() (bool, Shape) {
	if len(a[0]) < 2 {
		return false, a
	}

	b := a.Copy()

	x, y := b.zero()

	if x == 0 {
		return false, b
	}

	v := b[y][x-1]
	b[y][x-1] = 0
	b[y][x] = v

	return true, b
}

func (a Shape) Right() (bool, Shape) {
	if len(a[0]) < 2 {
		return false, a
	}

	b := a.Copy()

	x, y := b.zero()

	if x == len(a[0]) {
		return false, b
	}

	v := b[y][x+1]
	b[y][x+1] = 0
	b[y][x] = v

	return true, b
}

func (a Shape) Next() Moves {
	next := make(Moves, 0)

	// i, j := a.zero()
	// rows, columns := a.Rows(), a.Columns()

	return next
}

func LineCompare(e1, e2 Line) int {
	return slices.Compare(e1, e2)
}

func (e1 Shape) Equal(e2 Shape) int {
	return slices.CompareFunc(e1, e2, LineCompare)
}
