package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	inputs := loadInput()

	var (
		zeroCount = 0
		dial      = 50
	)
	for _, instruction := range inputs {
		dial = moveDial(dial, instruction)

		if dial == 0 {
			zeroCount++
		}
	}

	print("Number of zeros: ", zeroCount, "\n")
}

func moveDial(dial int, instruction string) int {
	steps, _ := strconv.Atoi(instruction[1:])
	if instruction[0] == 'R' {
		dial += steps
	} else {
		dial -= steps
	}

	if dial >= 100 {
		dial = dial % 100
	}

	for dial < 0 {
		dial = dial + 100
	}

	return dial
}

func loadInput() (inputs []string) {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	return
}
