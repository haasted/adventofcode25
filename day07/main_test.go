package main

import (
	"advent07/matrix"
	"advent07/textbuffer"
	"fmt"
	"math/big"
	"testing"
)

var testInput = `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`

func TestLogic(t *testing.T) {
	tb := textbuffer.New([]byte(testInput))
	splitCount, timelineCount := countSplits(tb)
	println(splitCount, timelineCount)
	assert(splitCount == 21)
	assert(timelineCount == 40)
}

func TestBig(t *testing.T) {
	bi := big.NewInt(0)
	bi.SetBit(bi, 5, 1)

	assert(fmt.Sprintf("%b", bi) == "100000")
	assert(bi.Bit(5) == 1)
	assert(bitCount(bi) == 1)
	assert(bi.Bit(209) == 0)
	assert(bi.BitLen() == 6)
}

func Test(t *testing.T) {
	buf := []byte(testInput)

	tb := textbuffer.New(buf)

	assert(tb.Get(0, 0) == '.')
	assert(tb.Get(7, 0) == 'S')
	assert(tb.Get(7, 2) == '^')

	assert(tb.IndexOf(0, 'S') == 7)
	assert(tb.IndexOf(0, 'X') == -1)
}

func Test2(t *testing.T) {
	buf := []byte("ABCDEFG")
	tb := textbuffer.New(buf)
	tb.Set(2, 0, 'X')

	line := tb.Line(0)
	assert(line == "ABXDEFG")
	println(line)

	tb.Set(6, 0, 'Z')
	line = tb.Line(0)
	assert(line == "ABXDEFZ")
}

func assert(in bool) {
	if !in {
		panic("assertion failed")
	}
}

func TestMatrix(t *testing.T) {
	m := matrix.Matrix{}

	m.Inc(5, 9, 1)
	m.Inc(1, 1, 2)
	println(m.String())
	assert(m.Get(1, 1) == 2)
	assert(m.Get(5, 9) == 1)

	m.Inc(15, 15, 1)
	for range 3 {
		m.Inc(2, 2, 4)
	}

	assert(m.Get(2, 2) == 12)

	println(m.String())
}
