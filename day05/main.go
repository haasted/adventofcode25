package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start, end int
}

func (r Range) isInRange(val int) bool {
	return r.start <= val && val <= r.end
}

type inventorySystem struct {
	ranges []Range
}

func (is *inventorySystem) addRange(r string) {
	is.ranges = append(is.ranges, newRange(r))
}

func (is *inventorySystem) isFresh(ingredient int) bool {
	for _, r := range is.ranges {
		if r.isInRange(ingredient) {
			return true
		}
	}
	return false
}

func newRange(r string) (res Range) {
	r = strings.TrimSpace(r)
	vals := strings.Split(r, "-")

	res.start, _ = strconv.Atoi(vals[0])
	res.end, _ = strconv.Atoi(vals[1])
	return
}

func main() {
	is, ingredients := loadFromInputFile()
	println("Ranges", len(is.ranges))
	println("Ingredients", len(ingredients))

	freshcount := 0
	for _, ingredient := range ingredients {
		if is.isFresh(ingredient) {
			freshcount++
		}
	}

	println("Number of fresh ingredients", freshcount)
}

func loadFromInputFile() (inventorySystem, []int) {
	file, _ := os.Open("./input")
	defer file.Close()
	return loadInput(file)
}

func loadInput(reader io.Reader) (is inventorySystem, ingredients []int) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			break
		}

		is.addRange(txt)
	}

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		ingredients = append(ingredients, i)
	}

	return
}
