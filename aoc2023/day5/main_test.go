package main

import (
	"log"
	"os"
	"testing"

	"github.com/10wpressure/advent-of-code/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	testFileName = "test.txt"
	testInput    = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, SolutionName())
}

var _ = Describe("Day 5", func() {
	BeforeEach(func() {
		_ = os.Remove(testFileName)
	})

	It("Part 1", func() {
		err := os.WriteFile(testFileName, []byte(testInput), 0666)
		if err != nil {
			log.Fatal(err)
		}
		f := util.OpenFile(testFileName)
		defer f.Close()
		defer os.Remove(testFileName)
		a := NewSolution()
		a.Parse(f)
		res := a.Part1()
		Expect(res).To(Equal("35"))
	})

	It("Part 2", func() {
		err := os.WriteFile(testFileName, []byte(testInput), 0666)
		if err != nil {
			log.Fatal(err)
		}
		f := util.OpenFile(testFileName)
		defer f.Close()
		defer os.Remove(testFileName)
		a := NewSolution()
		a.Parse(f)
		res := a.Part2()
		Expect(res).To(Equal("46"))
	})

})
