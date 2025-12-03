package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

// Test case from description
var inputs = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func TestPart1(t *testing.T) {
	var (
		dialValues         = []int{82, 52, 0, 95, 55, 0, 99, 0, 14, 32}
		expectedZeroCounts = []int{1, 0, 0, 0, 1, 0, 0, 0, 0, 1}
		dial               = 50
		zeroCount          int
		zeroPasses         int
	)

	scanner := bufio.NewScanner(strings.NewReader(inputs))
	for scanner.Scan() {
		var passes int
		dial, passes = moveDial(dial, scanner.Text())

		zeroPasses += passes
		if dialValues[0] != dial {
			panic(fmt.Sprint(dialValues[0], "!=", dial))
		}
		if expectedZeroCounts[0] != passes {
			println("dial value ", dial, passes, fmt.Sprintf(" %+v - %+v", dialValues, expectedZeroCounts))
			panic("unexpected zero count")
		}

		if dial == 0 {
			zeroCount++
		}

		dialValues = dialValues[1:]
		expectedZeroCounts = expectedZeroCounts[1:]
	}

	if zeroCount != 3 {
		println(zeroCount)
		panic("Unexpected amount of exact zeros")
	}

	if zeroCount+zeroPasses != 6 {
		println(zeroCount)
		panic("Unexpected amount of zero passes")
	}
}
