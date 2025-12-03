package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	inputs := loadInput()

	var (
		zeroCount  = 0
		zeroPasses = 0
		dial       = 50
	)
	for _, instruction := range inputs {
		var zerosPassed int
		dial, zerosPassed = moveDial(dial, instruction)

		if dial == 0 {
			zeroCount++
		}

		zeroPasses += zerosPassed
	}

	println("Number of zeros: ", zeroCount)
	println("Number of zeros passed: ", zeroPasses)
	println("Sum: ", zeroPasses+zeroCount)

}

func moveDial(dial int, instruction string) (int, int) {
	zeroCount := 0

	steps, _ := strconv.Atoi(instruction[1:])
	if instruction[0] == 'R' {
		dial += steps
	} else {
		if dial == 0 {
			zeroCount = -1
		}
		dial -= steps
	}

	for dial >= 100 {
		dial = dial - 100
		if dial != 0 {
			zeroCount++
		}
	}

	for dial < 0 {
		zeroCount++
		dial = dial + 100
	}

	return dial, zeroCount
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
