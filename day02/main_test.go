package main

import (
	"testing"
)

// Test ranges: "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
/*
   11-22 has two invalid IDs, 11 and 22.
   95-115 has one invalid ID, 99.
   998-1012 has one invalid ID, 1010.
   1188511880-1188511890 has one invalid ID, 1188511885.
   222220-222224 has one invalid ID, 222222.
   1698522-1698528 contains no invalid IDs.
   446443-446449 has one invalid ID, 446446.
   38593856-38593862 has one invalid ID, 38593859.
   The rest of the ranges contain no invalid IDs.
*/

func Test1(t *testing.T) {
	var IDs []string

	IDs = processRange("11-22")
	assertEqual(IDs, "11", "22")

	IDs = processRange("95-115")
	assertEqual(IDs, "99")

	IDs = processRange("1188511880-1188511890")
	assertEqual(IDs, "1188511885")

	IDs = processRange("38593856-38593862")
	assertEqual(IDs, "38593859")

	IDs = processRange("1698522-1698528")
	assertEqual(IDs)
}

func Test2(t *testing.T) {
	var IDs []string

	IDs = processRange2("11-22")
	assertEqual(IDs, "11", "22")

	IDs = processRange2("173-267")
	assertEqual(IDs, "222")

	IDs = processRange2("1188511880-1188511890")
	assertEqual(IDs, "1188511885")

	IDs = processRange2("38593856-38593862")
	assertEqual(IDs, "38593859")

	IDs = processRange2("1698522-1698528")
	assertEqual(IDs)

	IDs = processRange2("95-115")
	assertEqual(IDs, "99", "111")

	IDs = processRange2("998-1012")
	assertEqual(IDs, "999", "1010")

	IDs = processRange2("565653-565659")
	assertEqual(IDs, "565656")

	IDs = processRange2("1698522-1698528")
	assertEqual(IDs)

	IDs = processRange2("824824821-824824827")
	assertEqual(IDs, "824824824")

	IDs = processRange2("2121212118-2121212124")
	assertEqual(IDs, "2121212121")

	IDs = processRange2("2-20")
	assertEqual(IDs, "11")
}

func Test3(t *testing.T) {
	assertInvalid("123123123")
}

func assertInvalid(ID string) {
	if !isInvalid(ID) {
		println("expected to be invalid", ID)
		panic("")
	}
}

func assertEqual(IDs []string, expected ...string) {
	for _, id := range IDs {
		if id != expected[0] {
			println(id, " != ", expected[0])
			panic("")
		}

		expected = expected[1:]
	}

	if len(expected) != 0 {
		panic("More values expected")
	}
}
