package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/10wpressure/advent-of-code/util"
)

const (
	chunkSize = 100
)

type Point struct {
	x, y int64
}

func (p *Point) String() string {
	return fmt.Sprintf("{x:%d y:%d}", p.x, p.y)

}

func part1(f *os.File) string {
	var HouseMap = make(map[Point]int)

	fmt.Printf("########################################\n")
	var x, y int64 = 0, 0
	p := Point{0, 0}
	HouseMap[p] = 1

	util.ReadInChunks(f, chunkSize, func(buf []byte) bool {
		for _, v := range buf {
			switch v {
			case '^':
				y++
			case 'v':
				y--
			case '>':
				x++
			case '<':
				x--
			}
			p = Point{x: x, y: y}
			HouseMap[p]++
		}
		return false
	})
	cnt := 0
	HouseMap[p] = 1
	for k, v := range HouseMap {
		if v > 0 {
			cnt++
			fmt.Printf("%d: House %+v has been visited %d times\n", cnt, k, v)
		}
	}
	return strconv.Itoa(cnt)
}

func part2() string {
	return fmt.Sprintf("")
}

func main() {
	f := util.OpenFile("input.txt")
	defer f.Close()

	res1 := part1(f)
	log.Println("Part 1:", res1)

	res2 := part2()
	log.Println("Part 2:", res2)
}
