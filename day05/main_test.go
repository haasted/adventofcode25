package main

import (
	"strings"
	"testing"
)

func Test(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32
`

	is, ingredients := loadInput(strings.NewReader(input))
	assert(len(is.ranges) == 4)
	assert(len(ingredients) == 6)

	for _, ingredient := range []int{1, 8, 32} {
		assert(!is.isFresh(ingredient))
	}

	for _, ingredient := range []int{5, 11, 17} {
		assert(is.isFresh(ingredient))
	}
}

func assert(in bool) {
	if !in {
		panic("assertion failed")
	}
}
