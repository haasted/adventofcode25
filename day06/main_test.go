package main

import (
	"slices"
	"strings"
	"testing"
)

const input = `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +
`

func Test(t *testing.T) {
	calcs := loadInput(strings.NewReader(input))

	assert(len(calcs) == 4)
	assert(calcs[0].operator == "*")
	assert(len(calcs[0].operands) == 3)
	assert(slices.Equal(calcs[0].operands, []int{123, 45, 6}))

	assert(calcs[0].run() == 33210)
	assert(calcs[1].run() == 490)
	assert(calcs[2].run() == 4243455)
	assert(calcs[3].run() == 401)
}

func TestSplitter(t *testing.T) {
	input := "  5     7  "
	input = strings.TrimSpace(input)
	split := wsSplitter.Split(input, -1)
	for i, v := range split {
		println(i, " ", v)
	}
}

func assert(in bool) {
	if !in {
		panic("assertion failed")
	}
}
