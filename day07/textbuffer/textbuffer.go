package textbuffer

import (
	"bufio"
	"bytes"
	"fmt"
)

type TextBuffer struct {
	lines []string
}

func New(in []byte) *TextBuffer {
	var lines []string

	scanner := bufio.NewScanner(bytes.NewReader(in))
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return &TextBuffer{lines}
}

// No fancy Rune math here. Everything is ascii
func (tb *TextBuffer) Get(x, y int) byte {
	tb.checkUpperBounds(x, y)

	return tb.lines[y][x]
}

func (tb *TextBuffer) Set(x, y int, value rune) {
	tb.checkUpperBounds(x, y)

	line := tb.lines[y]
	line = fmt.Sprintf("%s%s%s", line[:x], string(value), line[x+1:])
	tb.lines[y] = line
}

func (tb *TextBuffer) IndexOf(line int, char rune) (idx int) {
	tb.checkUpperBounds(-1, line)

	var c rune
	for idx, c = range tb.lines[line] {
		if c == char {
			return
		}
	}

	return -1
}

func (tb *TextBuffer) LineCount() int {
	return len(tb.lines)
}

func (tb *TextBuffer) Line(line int) string {
	tb.checkUpperBounds(-1, line)

	return tb.lines[line]
}

func (tb *TextBuffer) checkUpperBounds(x, y int) {
	if y >= len(tb.lines) {
		panic("You fool!")
	}

	if x >= len(tb.lines[y]) {
		panic("You utter foool!")
	}
}
