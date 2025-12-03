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

func Test1(t *testing.T) {
	var (
		dialValues = []int{82, 52, 0, 95, 55, 0, 99, 0, 14, 32}
		dial       = 50
	)

	scanner := bufio.NewScanner(strings.NewReader(inputs))
	for scanner.Scan() {
		dial = moveDial(dial, scanner.Text())
		if dialValues[0] != dial {
			panic(fmt.Sprint(dialValues[0], "!=", dial))
		}
		dialValues = dialValues[1:]
	}
}
