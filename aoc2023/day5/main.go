package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	t "github.com/10wpressure/advent-of-code/types"
	"github.com/10wpressure/advent-of-code/util"
)

const (
	inputFile = "input.txt"
	chunkSize = 100
)

type Solution struct {
	seeds   []int64
	mapping []*Mapping
}

type Mapping struct {
	maps []*SingleMap
}

type SingleMap struct {
	Range t.Range
	Delta int64
}

func NewMapping() *Mapping {
	return &Mapping{}
}

func (m *Mapping) AddMapping(dest, src, len int64) {
	m.maps = append(m.maps, &SingleMap{
		Range: t.Range{
			Start: src,
			End:   src + len,
		},
		Delta: dest - src,
	})
}

func (m *Mapping) ReverseLookup(val int64) int64 {
	for _, cur := range m.maps {
		rev := val - cur.Delta
		if cur.Range.Contains(rev) {
			return rev
		}
	}
	return val
}

func (m *Mapping) ApplyMap(val int64) int64 {
	for _, cur := range m.maps {
		if cur.Range.Contains(val) {
			return val + cur.Delta
		}
	}
	return val
}

func NewSolution() *Solution {
	return &Solution{}
}

func (s *Solution) SeedRanges() []t.Range {
	var seedRanges = make([]t.Range, 0)
	for i := 0; i < len(s.seeds); i += 2 {
		var current = s.seeds[i]
		var delta = s.seeds[i+1]

		seedRanges = append(seedRanges, t.Range{
			Start: current,
			End:   current + delta,
		})
	}
	return seedRanges
}

func (s *Solution) Parse(f *os.File) {
	scanner := bufio.NewScanner(f)
	scanner.Scan()

	seedsLine := scanner.Text()
	stringOfNumbers := strings.Fields(seedsLine)[1:]
	for _, numStr := range stringOfNumbers {
		num, _ := strconv.Atoi(numStr)
		s.seeds = append(s.seeds, int64(num))
	}

	curMap := NewMapping()
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.Contains(line, "map:") {
			s.mapping = append(s.mapping, curMap)
			curMap = NewMapping()
			continue
		}

		l := strings.Fields(line)
		nums := make([]int64, 3)
		for i, numStr := range l {
			num, _ := strconv.Atoi(numStr)
			nums[i] = int64(num)
		}
		curMap.AddMapping(nums[0], nums[1], nums[2])
	}
	if len(curMap.maps) != 0 {
		s.mapping = append(s.mapping, curMap)
	}
}

func (s *Solution) Part1() string {
	var minimum int64 = math.MaxInt64
	for _, seed := range s.seeds {
		var current = seed
		for _, m := range s.mapping {
			current = m.ApplyMap(current)
		}
		minimum = min(minimum, current)
	}

	return strconv.FormatInt(minimum, 10)
}

func (s *Solution) Part2() string {
	ranges := s.SeedRanges()

	var location int64 = 1
	var reversed = s.mapping
	slices.Reverse(reversed)
	for {
		cur := location
		for _, m := range reversed {
			cur = m.ReverseLookup(cur)
		}
		for _, sr := range ranges {
			if sr.Contains(cur) {
				return strconv.FormatInt(location, 10)
			}
		}
		location++
	}

}

func main() {
	f := util.OpenFile(inputFile)
	defer f.Close()
	a := NewSolution()
	a.Parse(f)
	part1 := a.Part1()
	part2 := a.Part2()

	log.Printf("Part 1: %s", part1)
	log.Printf("Part 2: %s", part2)
}
