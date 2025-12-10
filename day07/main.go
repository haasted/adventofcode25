package main

import (
	"advent07/matrix"
	"advent07/textbuffer"
	"io"
	"math/big"
	"os"
)

func main() {
	buf := getInput()
	tb := textbuffer.New(buf)

	splitCount, timelineCount := countSplits(tb)

	println("Split count", splitCount)
	if splitCount != 1539 {
		println("WARN! Error in split calculation somewhere")
	}

	println("Timeline count", timelineCount)
	if timelineCount != 6479180385864 {
		println("WARN! Error in timelines calculation somewhere")
	}
}

func countSplits(tb *textbuffer.TextBuffer) (splitCount, timelineCount int) {
	startIdx := tb.IndexOf(0, 'S')
	m := matrix.Matrix{}

	m.Inc(startIdx, 0, 1) // Set source

	for lineidx := range tb.LineCount() - 1 {
		for x := range m.RowLen(lineidx) {
			val := m.Get(x, lineidx)
			if val == 0 {
				continue
			}

			nextLine := 1 + lineidx // readability
			if tb.Get(x, nextLine) == '^' {
				splitCount++

				m.Inc(x-1, nextLine, val)
				m.Inc(x+1, nextLine, val)
			} else {
				m.Inc(x, nextLine, val)
			}
		}
	}

	println(m.String())
	sum := 0
	m.IterLastRow(func(val int) {
		sum += val
	})

	return splitCount, sum
}

func bitCount(bi *big.Int) (count uint) {
	for idx := range bi.BitLen() {
		count += bi.Bit(idx)
	}
	return
}

func getInput() []byte {
	file, _ := os.Open("./input")
	buf, _ := io.ReadAll(file)
	return buf
}
