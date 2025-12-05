package main

import (
	"bufio"
	"os"
)

type cellState int

const (
	empty cellState = iota
	containsRoll
	rollMarkedForRemoval
)

type board struct {
	layout [][]cellState
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

	switch b.layout[y][x] {
	case containsRoll, rollMarkedForRemoval:
		return true
	default:
		return false
	}
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

func (b *board) markAllAccessible() (count int) {
	for y := range b.layout {
		for x := range b.layout[y] {
			if b.isRoll(x, y) && b.isAccessible(x, y) {
				count++
				b.layout[y][x] = rollMarkedForRemoval
			}
		}
	}

	return
}

func (b *board) removeAllMarked() {
	for y := range b.layout {
		for x := range b.layout[y] {
			if b.layout[y][x] == rollMarkedForRemoval {
				b.layout[y][x] = empty
			}
		}
	}
}

func (b *board) addRow(row string) {
	r := make([]cellState, len(row))
	for i, c := range row {
		if c == '@' {
			r[i] = containsRoll
		}
	}

	b.layout = append(b.layout, r)
}

func main() {
	b := loadInput()
	count := b.markAllAccessible()

	println("Accessible rolls: ", count)

	b.removeAllMarked()
	marked := count
	for marked > 0 {
		marked = b.markAllAccessible()
		b.removeAllMarked()
		count += marked
	}

	println("Total removed", count)
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
