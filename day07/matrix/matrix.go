package matrix

import (
	"fmt"
	"slices"
	"strings"
)

type Matrix struct {
	contents [][]int
}

func (m *Matrix) Inc(x, y, val int) {
	m.ensureSize(x, y)
	m.contents[y][x] = m.contents[y][x] + val
}

func (m *Matrix) RowLen(y int) int {
	m.ensureSize(0, y)
	return len(m.contents[y])
}

func (m *Matrix) IterLastRow(it func(int)) {
	y := len(m.contents)

	for _, val := range m.contents[y-1] {
		it(val)
	}
}

func (m *Matrix) RowSum(y int) (sum int) {
	for _, val := range m.contents[y] {
		sum += val
	}

	return
}

func (m *Matrix) Get(x, y int) int {
	m.ensureSize(x, y)
	return m.contents[y][x]
}

func (m *Matrix) ensureSize(x, y int) {
	if y >= len(m.contents) {
		m.contents = slices.Grow(m.contents, y+1)
		m.contents = m.contents[:y+1]
	}

	if x >= len(m.contents[y]) {
		m.contents[y] = slices.Grow(m.contents[y], x+1)
		m.contents[y] = m.contents[y][:x+1]
	}
}

func (m *Matrix) String() string {
	var sb strings.Builder
	for _, row := range m.contents {
		for _, val := range row {
			sb.WriteString(fmt.Sprintf("%v ", val))
		}

		sb.WriteRune('\n')
	}

	return sb.String()
}
