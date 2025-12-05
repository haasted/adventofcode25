package main

import (
	"bufio"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	start, end int
}

func (r Range) isInRange(val int) bool {
	if val == 138028264431749 {
		if r.start <= val && val <= r.end {
			println("in range", r.start, r.end)
		}
	}

	return r.start <= val && val <= r.end
}

func (r Range) size() int {
	return 1 + r.end - r.start
}

type inventorySystem struct {
	ranges []Range
}

func (is *inventorySystem) addRange(r string) {
	is.ranges = append(is.ranges, newRange(r))
}

func (is *inventorySystem) sortAndMergeRanges() {
	slices.SortFunc(is.ranges, func(a, b Range) int {
		return a.start - b.start
	})

	merged := []Range{is.ranges[0]} // Ignore empty ranges.
	for _, r := range is.ranges {
		highest := &merged[len(merged)-1]
		if highest.end >= r.start {
			highest.end = max(r.end, highest.end)
			continue
		}

		merged = append(merged, r)
	}

	is.ranges = merged
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

	var err error
	res.start, err = strconv.Atoi(vals[0])
	if err != nil {
		panic(err)
	}

	res.end, err = strconv.Atoi(vals[1])
	if err != nil {
		panic(err)
	}

	return
}

func main() {
	is, ingredients := loadFromInputFile()
	is.sortAndMergeRanges()

	println("Ranges", len(is.ranges))

	println("Ingredients", len(ingredients))

	fresh1 := []int{}
	freshcount := 0
	for _, ingredient := range ingredients {
		if is.isFresh(ingredient) {
			fresh1 = append(fresh1, ingredient)
			freshcount++
		}
	}

	println("Number of fresh ingredients", freshcount)

	totalFreshIngredientCount := 0
	for _, r := range is.ranges {
		totalFreshIngredientCount += r.size()
	}

	println("Number of possibly fresh ingredients", totalFreshIngredientCount)

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
