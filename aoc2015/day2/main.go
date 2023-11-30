package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Box struct {
	l, w, h int64
}

func NewBox(l, w, h int64) *Box {
	return &Box{l, w, h}
}

func (b *Box) SurfaceArea() int64 {
	return 2*b.l*b.w + 2*b.w*b.h + 2*b.h*b.l
}

func (b *Box) ExtraSurfaceArea() int64 {
	return min(b.l*b.w, b.w*b.h, b.h*b.l)
}

func (b *Box) SurfaceAreaWithExtra() int64 {
	return b.SurfaceArea() + b.ExtraSurfaceArea()
}

func (b *Box) ASCSortSides() []int64 {
	sortedBoxSides := []int64{b.l, b.w, b.h}
	sort.Slice(sortedBoxSides, func(i, j int) bool {
		return sortedBoxSides[i] < sortedBoxSides[j]
	})
	return sortedBoxSides
}
func (b *Box) CalculateRibbon() int64 {
	sides := b.ASCSortSides()
	wrap := sides[0]*2 + sides[1]*2
	bow := b.l * b.w * b.h

	return bow + wrap
}

func part1() string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	scanner := bufio.NewScanner(f)

	var wrappingPaper int64 = 0
	for scanner.Scan() {
		line := scanner.Text()

		box := NewBox(0, 0, 0)
		_, err := fmt.Sscanf(line, "%dx%dx%d", &box.l, &box.w, &box.h)
		if err != nil {
			log.Fatal(err)
		}

		wrappingPaper += box.SurfaceAreaWithExtra()
	}

	return fmt.Sprintf("%d", wrappingPaper)
}

func part2() string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	scanner := bufio.NewScanner(f)

	var ribbon int64 = 0
	for scanner.Scan() {
		line := scanner.Text()

		box := NewBox(0, 0, 0)
		_, err := fmt.Sscanf(line, "%dx%dx%d", &box.l, &box.w, &box.h)
		if err != nil {
			log.Fatal(err)
		}

		ribbon += box.CalculateRibbon()
	}

	return fmt.Sprintf("%d", ribbon)
}

func main() {
	res1 := part1()
	log.Println("Part 1:", res1)

	res2 := part2()
	log.Println("Part 2:", res2)

}
