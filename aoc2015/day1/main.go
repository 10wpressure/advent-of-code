package main

import (
	"fmt"
	"log"
	"os"

	"github.com/10wpressure/advent-of-code/util"
)

const (
	chunkSize = 100
)

func part1(f *os.File) string {
	floor := 0

	util.ReadInChunks(f, chunkSize, func(buf []byte) bool {
		for i := range buf {
			switch string(buf[i]) {
			case "(":
				floor++
			case ")":
				floor--
			}
		}
		return false
	})

	return fmt.Sprintf("%d", floor)
}

func part2(f *os.File) string {
	floor := 0
	pos := 0

	util.ReadInChunks(f, chunkSize, func(buf []byte) bool {
		for i := 0; i < len(buf); i++ {
			switch string(buf[i]) {
			case "(":
				floor++
			case ")":
				floor--
			}
			pos++
			if floor < 0 {
				return true
			}
		}
		return false
	})

	return fmt.Sprintf("%d", pos)
}

func main() {
	f := util.OpenFile("input.txt")
	defer f.Close()
	res1 := part1(f)
	log.Println("Part 1:", res1)

	res2 := part2(f)
	log.Println("Part 2:", res2)
}
