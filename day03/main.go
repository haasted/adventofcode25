package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type bank []rune

func (b bank) findLargest(startIdx, endIdx int) (idx int, value rune) {
	for i := startIdx; i <= endIdx; i++ {
		if b[i] > value {
			value, idx = b[i], i
		}
	}

	return
}

func (b bank) findMaxJoltage(digits int) int {
	var (
		result = make([]rune, digits)
		idx    = -1
	)

	for digit := range digits {
		idx, result[digit] = b.findLargest(idx+1, len(b)-(digits-digit))
	}

	var bldr strings.Builder
	for _, r := range result {
		bldr.WriteRune(r + '0')
	}

	res, err := strconv.Atoi(bldr.String())
	if err != nil {
		panic(err)
	}
	return res
}

func main() {
	var sum int
	joltages := loadInput()
	for _, j := range joltages {
		joltage := j.findMaxJoltage(2)
		sum += joltage
	}

	println("Part1 total joltage", sum)

	sum = 0
	for _, j := range joltages {
		joltage := j.findMaxJoltage(12)
		sum += joltage
	}

	println("Part2 total joltage", sum)
}

func loadInput() (inputs []bank) {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, newBank(scanner.Text()))
	}

	return
}

func newBank(in string) (b bank) {
	in = strings.TrimSpace(in)
	for _, v := range in {
		b = append(b, v-'0')
	}

	return
}
