package main

import (
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/10wpressure/advent-of-code/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	testFileName = "test.txt"
	testInput    = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, SolutionName())
}

var _ = Describe("Day 3", func() {
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
		res := part1(f)
		Expect(res).To(Equal("4361"))
	})

	It("Part 2", func() {
		err := os.WriteFile(testFileName, []byte(testInput), 0666)
		if err != nil {
			log.Fatal(err)
		}
		f := util.OpenFile(testFileName)
		defer f.Close()
		defer os.Remove(testFileName)
		res := part2(f)
		s := 467835
		Expect(res).To(Equal(strconv.Itoa(s)))
	})

})
