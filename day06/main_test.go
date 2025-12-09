package main

import (
	"fmt"
	"slices"
	"strings"
	"testing"
	"unicode"
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
	assert(slices.Equal(calcs[0].operands, []string{"123", "45", "6"}))

	assert(calcs[0].run() == 33210)
	assert(calcs[1].run() == 490)
	assert(calcs[2].run() == 4243455)
	assert(calcs[3].run() == 401)

	calcs[0].c8dize()
	assert(calcs[0].run() == 8544)

	calcs[1].c8dize()
	assert(calcs[1].run() == 3253600)

	calcs[2].c8dize()
	println(calcs[2].run())
	calcs[3].c8dize()
	println(calcs[3].run())

}

func TestTransposeStrings(t *testing.T) {
	input := []string{"328", "64", "98"}
	// result: 356 * 24 * 1

	maxLength := 0
	for _, i := range input {
		maxLength = max(len(i), maxLength)
	}

	for idx := range len(input) {
		input[idx] = fmt.Sprintf("%*s", maxLength, input[idx])
	}

	fmt.Printf("%#v\n", input)

	results := make([]strings.Builder, maxLength)
	for idx := maxLength - 1; idx >= 0; idx-- {
		for _, s := range input {
			r := rune(s[idx])
			if unicode.IsDigit(r) {
				results[idx].WriteRune(r)
			}
		}
	}

	// TODO : Use the operator to find the beginning of the numbers....

	for _, b := range results {
		println(b.String())
	}
}

func TestParseCephalopod(t *testing.T) {
	input := `
123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +  `

	println(input)

}

func assert(in bool) {
	if !in {
		panic("assertion failed")
	}
}
