package main

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var wsSplitter = regexp.MustCompile(`\s+`)

type calculation struct {
	operator string
	operands []int
}

func (c calculation) run() (res int) {
	switch c.operator {
	case "*":
		res = 1
		for _, val := range c.operands {
			res *= val
		}
	case "+":
		for _, val := range c.operands {
			res += val
		}
	default:
		println("Unknown operator", c.operator)
		panic("")
	}

	return
}

func main() {
	calcs := loadFromInputFile()

	sum := 0
	for _, c := range calcs {
		sum += c.run()
	}

	println("Total sum of all calculations", sum)

}

func loadFromInputFile() []calculation {
	file, _ := os.Open("./input")
	defer file.Close()
	return loadInput(file)
}

func loadInput(reader io.Reader) (calcs []calculation) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		cols := wsSplitter.Split(line, -1)

		if len(calcs) < len(cols) {
			calcs = make([]calculation, len(cols))
		}

		if unicode.IsDigit(rune(line[0])) {
			for idx, val := range cols {
				i, err := strconv.Atoi(val)
				if err != nil {
					panic(err)
				}
				calcs[idx].operands = append(calcs[idx].operands, i)
			}
		} else {
			for idx, op := range cols {
				calcs[idx].operator = op
			}
		}
	}

	return
}
