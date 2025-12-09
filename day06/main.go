package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type calculation struct {
	operator string
	operands []string
}

// celaphopodize the calculation
func (c *calculation) c8dize() {
	input := c.operands // Clone?

	maxLength := 0
	for _, i := range input {
		maxLength = max(len(i), maxLength)
	}

	for idx := range len(input) {
		input[idx] = fmt.Sprintf("%*s", maxLength, input[idx])
	}

	results := make([]strings.Builder, maxLength)
	for idx := maxLength - 1; idx >= 0; idx-- {
		for _, s := range input {
			r := rune(s[idx])
			if unicode.IsDigit(r) {
				results[idx].WriteRune(r)
			}
		}
	}

	c.operands = nil

	for _, r := range results {
		c.operands = append(c.operands, r.String())
	}
}

func (c calculation) run() (res int) {
	switch c.operator {
	case "*":
		res = 1
		for _, val := range c.operands {
			i, _ := strconv.Atoi(val)
			res *= i
		}
	case "+":
		for _, val := range c.operands {
			i, _ := strconv.Atoi(val)
			res += i
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
	if sum != 5346286649122 {
		println(" !! WARN !! - Currently a wrong result!")
	}

	sum = 0
	for _, c := range calcs {
		c.c8dize()
		sum += c.run()
	}

	println("Total sum of all celaphopod calculations", sum)

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
		cols := strings.Fields(line)

		if len(calcs) < len(cols) {
			calcs = make([]calculation, len(cols))
		}

		if unicode.IsDigit(rune(line[0])) {
			for idx, val := range cols {
				calcs[idx].operands = append(calcs[idx].operands, val)
			}
		} else {
			for idx, op := range cols {
				calcs[idx].operator = op
			}
		}
	}

	return
}
