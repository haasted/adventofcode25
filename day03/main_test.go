package main

import "testing"

func Test1(t *testing.T) {
	b := newBank("0123456")
	for i := range b {
		assert(b[i] == rune(i))
	}

	b = newBank("456")
	assert(b[0] == rune(4))
	assert(b[1] == rune(5))
	assert(b[2] == rune(6))
}

func Test2(t *testing.T) {
	b := newBank("2141132231122222128231324232242223212124232222531422221312221242422234223122222212222542422344222321")
	idx, value := b.findLargest(0, len(b))

	assert(idx == 18)
	assert(value == 8)

	idx, value = b.findLargest(idx+1, len(b))
	assert(value == 5)
	assert(idx == 46)

}

func Test3(t *testing.T) {
	var mj int
	mj = newBank("987654321111111").findMaxJoltage(2)
	assert(mj == 98)
	mj = newBank("811111111111119").findMaxJoltage(2)
	assert(mj == 89)
	mj = newBank("234234234234278").findMaxJoltage(2)
	assert(mj == 78)
	mj = newBank("818181911112111").findMaxJoltage(2)
	assert(mj == 92)
}

func Test4(t *testing.T) {
	var mj int

	mj = newBank("987654321111111").findMaxJoltage(12)
	assert(mj == 987654321111)

	mj = newBank("811111111111119").findMaxJoltage(12)
	assert(mj == 811111111119)

	mj = newBank("234234234234278").findMaxJoltage(12)
	assert(mj == 434234234278)

	mj = newBank("818181911112111").findMaxJoltage(12)
	assert(mj == 888911112111)
}

func assert(in bool) {
	if !in {
		panic("assertion failed")
	}
}
