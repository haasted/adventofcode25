package main

import (
	"bufio"
	"os"
	"strings"
)

type bank []rune

func (b bank) findLargest(startIdx, endIdx int) (idx int, value rune) {
	for i := startIdx; i < endIdx; i++ {
		if b[i] > value {
			value, idx = b[i], i
		}
	}

	return
}

func (b bank) findMaxJoltage() int {
	idx, d1 := b.findLargest(0, len(b)-1)
	_, d2 := b.findLargest(idx+1, len(b))

	return int(d1*10) + int(d2)
}

func main() {
	var sum int
	joltages := loadInput()
	for _, j := range joltages {
		joltage := j.findMaxJoltage()
		sum += joltage
	}

	println("Totalt joltage", sum)
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
