package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alendavid/puzzling/pkg/shape"
	"github.com/alendavid/puzzling/pkg/solver"
)

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	handleErr(err)

	line := strings.Split(strings.TrimSpace(text), " ")
	if len(line) != 2 {
		handleErr(fmt.Errorf("input: [rows] [columns]"))
	}

	rows, err := strconv.Atoi(line[0])
	handleErr(err)

	columns, err := strconv.Atoi(line[1])
	handleErr(err)

	puzzle := make(shape.Shape, rows)

	for i := 0; i < rows; i += 1 {
		text, err := reader.ReadString('\n')
		handleErr(err)

		line := strings.Split(strings.TrimSpace(text), " ")

		if len(line) != columns {
			handleErr(fmt.Errorf("declared column size as %d, received %d", columns, len(line)))
		}

		puzzle[i] = make(shape.Line, columns)

		for j := 0; j < columns; j += 1 {
			val, err := strconv.Atoi(line[j])
			handleErr(err)

			puzzle[i][j] = val
		}
	}

	fmt.Println(solver.Solve(puzzle))
}
