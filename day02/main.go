package main

import (
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	ranges := loadInput()

	sum := 0
	for _, r := range ranges {
		r = strings.TrimSpace(r)

		invalidIDs := processRange2(r)
		for _, id := range invalidIDs {
			// println("Invalid ID ", id)
			i, err := strconv.Atoi(id)
			if err != nil {
				panic(err)
			}
			sum += i
		}
	}

	println("Final sum", sum)
}

func processRange2(r string) (invalidIDs []string) {
	start, end := split(r)

	for v := start; v <= end; v++ {
		value := strconv.Itoa(v)
		if isInvalid(value) {
			invalidIDs = append(invalidIDs, value)
			continue
		}
	}

	return
}

func isInvalid(ID string) bool {
	if len(ID) == 1 { // Note range 2-20 ;)
		return false
	}

	half := len(ID) / 2
	half = half + len(ID)%2

	for size := 1; size <= half; size++ {
		slice := ID[:size]

		// TODO use strings.Repeat
		var bldr strings.Builder
		for bldr.Len() < len(ID) {
			bldr.WriteString(slice)
		}
		test := bldr.String()

		if test == ID {
			return true
		}
	}

	return false
}

func processRange(r string) (invalidIDs []string) {
	start, end := split(r)
	for v := start; v <= end; v++ {
		value := strconv.Itoa(v)
		if len(value)%2 == 1 {
			continue
		}

		split := len(value) >> 1
		if value[:split] == value[split:] {
			invalidIDs = append(invalidIDs, value)
		}
	}

	return
}

func split(r string) (start, end int) {
	vals := strings.Split(r, "-")
	start, _ = strconv.Atoi(vals[0])
	end, _ = strconv.Atoi(vals[1])
	return
}

func loadInput() (ranges []string) {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	input, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	ranges = strings.Split(string(input), ",")
	return
}
