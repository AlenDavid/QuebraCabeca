package shape

import (
	"fmt"
	"math"
	"slices"
)

type Line []int
type Shape []Line
type Moves []Shape

func (s Shape) Pos(value int) (i, j int) {
	for i = 0; i < len(s); i += 1 {
		for j = 0; j < len(s[i]); j += 1 {
			if s[i][j] == value {
				return
			}
		}
	}

	return -1, -1
}

func (s Shape) String() string {
	t := "[ "

	for i, line := range s {
		for j, n := range line {
			t += fmt.Sprintf("%d", n)

			if j*i < len(line) {
				t += ", "
			}
		}
	}

	return t + " ]"
}

func (m Moves) String() string {
	t := ""

	for _, shape := range m {
		t += fmt.Sprintf("%s\n", shape.String())
	}

	return t
}

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

func Final() Shape {
	return Shape{
		{1, 2, 3},
		{8, 0, 4},
		{7, 6, 5},
	}
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

	if y == len(a)-1 {
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

	if x == len(a[0])-1 {
		return false, b
	}

	v := b[y][x+1]
	b[y][x+1] = 0
	b[y][x] = v

	return true, b
}

func (a Shape) Next() Moves {
	next := make(Moves, 0)

	if moving, shape := a.Up(); moving {
		next = append(next, shape)
	}

	if moving, shape := a.Right(); moving {
		next = append(next, shape)
	}

	if moving, shape := a.Down(); moving {
		next = append(next, shape)
	}

	if moving, shape := a.Left(); moving {
		next = append(next, shape)
	}

	return next
}

func LineCompare(e1, e2 Line) int {
	return slices.Compare(e1, e2)
}

func (e1 Shape) Distance(e2 Shape) (distance int) {
	for i, line := range e2 {
		for j := range line {
			value := e2[i][j]
			x, y := e1.Pos(value)

			distance += int(math.Pow(float64(x-i), 2) + math.Pow(float64(j-y), 2))
		}
	}

	return
}

func (e1 Shape) Equal(e2 Shape) bool {
	return slices.CompareFunc(e1, e2, LineCompare) == 0
}

func (m Moves) Contains(a Shape) bool {
	return slices.ContainsFunc(m, func(b Shape) bool {
		return a.Equal(b)
	})
}

func (m Moves) Sort() {
	final := Final()

	slices.SortFunc(m, func(a, b Shape) int {
		return a.Distance(final) - b.Distance(final)
	})
}
