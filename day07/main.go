package main

import (
	"advent07/textbuffer"
	"fmt"
	"io"
	"math/big"
	"os"
)

func main() {
	buf := getInput()
	tb := textbuffer.New(buf)
	columnCount := len(tb.Line(0))

	startIdx := tb.IndexOf(0, 'S')

	rays := big.NewInt(0)
	rays.SetBit(rays, startIdx, 1)

	splitCount := 0
	for lineidx := range tb.LineCount() - 1 {
		fmt.Printf("%0*b - %d\n", columnCount, rays, bitCount(rays))
		nextRays := big.NewInt(0)

		for bit := range rays.BitLen() {
			if rays.Bit(bit) == 0 {
				continue
			}

			if tb.Get(bit, lineidx+1) == '^' {
				splitCount++
				nextRays.SetBit(nextRays, bit-1, 1)
				nextRays.SetBit(nextRays, bit+1, 1)
			} else {
				nextRays.SetBit(nextRays, bit, 1)

			}
		}

		rays = nextRays
	}

	println("Split count", splitCount)
}

func bitCount(bi *big.Int) (count uint) {
	for idx := range bi.BitLen() {
		count += bi.Bit(idx)
	}
	return
}

func getInput() []byte {
	file, _ := os.Open("./input")
	buf, _ := io.ReadAll(file)
	return buf
}
