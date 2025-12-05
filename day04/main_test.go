package main

import (
	"bufio"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	b := board{}

	b.addRow("..@@..@@")
	b.addRow("@......@")

	assert(!b.isRoll(0, 0))
	assert(b.isRoll(2, 0))

	assert(!b.isRoll(-1, 2))
	assert(!b.isRoll(5, -1))
	assert(!b.isRoll(100, 0))
	assert(!b.isRoll(2, 100))
}

func Test2(t *testing.T) {
	input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	b := board{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		b.addRow(scanner.Text())
	}

	assert(!b.isRoll(0, 0))
	assert(b.isRoll(2, 0))
	assert(b.isRoll(0, 9))
	assert(!b.isRoll(9, 9))

	count := 0
	for y := range b.layout {
		for x := range b.layout[y] {
			if b.isRoll(x, y) && b.isAccessible(x, y) {
				count++
			}
		}
	}

	assert(count == 13)
}

func Test3(t *testing.T) {
	b := board{}
	b.addRow("..@")
	b.addRow("@@.")
	b.addRow(".@@")

	assert(!b.isAccessible(1, 1))
	assert(b.isAccessible(0, 1))
}

func assert(in bool) {
	if !in {
		panic("assertion failed")
	}
}
