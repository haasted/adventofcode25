package main

import (
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
	// file, _ := os.Open("./input")
	// buf, _ := io.ReadAll(file)
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
