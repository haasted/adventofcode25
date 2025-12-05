package main

import (
	"bufio"
	"os"
)

type board struct {
	layout [][]bool
}

func (b *board) isRoll(x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}

	if y >= len(b.layout) {
		return false
	}

	if x >= len(b.layout[y]) {
		return false
	}

	return b.layout[y][x]
}

func (b *board) isAccessible(x, y int) bool {
	count := -1

	for xd := range 3 {
		xd = xd - 1
		for yd := range 3 {
			yd = yd - 1
			if b.isRoll(x+xd, y+yd) {
				count++
			}
		}
	}

	return count < 4
}

func (b *board) addRow(row string) {
	r := make([]bool, len(row))
	for i, c := range row {
		r[i] = c == '@'
	}

	b.layout = append(b.layout, r)
}

func main() {
	b := loadInput()
	count := 0
	for y := range b.layout {
		for x := range b.layout[y] {
			if b.isRoll(x, y) && b.isAccessible(x, y) {
				count++
			}
		}
	}

	println("Accessible rows: ", count)
}

func loadInput() (b board) {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		b.addRow(scanner.Text())
	}

	return
}
