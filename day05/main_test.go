package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32
`

	is, ingredients := loadInput(strings.NewReader(input))
	assert(len(is.ranges) == 4)
	assert(len(ingredients) == 6)

	is.sortAndMergeRanges()
	assert(len(is.ranges) == 2)
	assert(is.ranges[0].start == 3)
	assert(is.ranges[0].end == 5)
	assert(is.ranges[1].start == 10)
	assert(is.ranges[1].end == 20)

	for _, ingredient := range []int{1, 8, 32} {
		assert(!is.isFresh(ingredient))
	}

	for _, ingredient := range []int{5, 11, 17} {
		assert(is.isFresh(ingredient))
	}

	count := 0
	for _, r := range is.ranges {
		count += r.size()
	}

	assert(count == 14)
}

func TestSortAndMerge1(t *testing.T) {
	is := inventorySystem{}
	is.addRange("5-10")
	is.addRange("3-8")
	assert(is.ranges[0].start == 5)
	assert(is.ranges[1].start == 3)
	is.sortAndMergeRanges()
	assert(is.ranges[0].start == 3)
	assert(is.ranges[0].end == 10)
}
func TestSortAndMerge2(t *testing.T) {
	is := inventorySystem{}
	is.addRange("5-10")
	is.addRange("5-7")

	is.addRange("20-22")
	is.addRange("15-20")

	is.addRange("29-29")

	is.sortAndMergeRanges()

	fmt.Printf("%+v\n", is.ranges)

	println(is.isFresh(29))
	println(is.isFresh(27))
}

func TestSortAndMerge234234(t *testing.T) {
	is := inventorySystem{}

	is.addRange("302553299774028-302939011277575")
	is.addRange("72206822427466-79146949214971")
	is.addRange("302939011277575-303124699217159")

	is.sortAndMergeRanges()
	assert(len(is.ranges) == 2)
	assert(is.ranges[0].start == 72206822427466)
	assert(is.ranges[0].end == 79146949214971)

	assert(is.ranges[1].start == 302553299774028)
	assert(is.ranges[1].end == 303124699217159)
}

func assert(in bool) {
	if !in {
		panic("assertion failed")
	}
}
