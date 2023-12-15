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
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, SolutionName())
}

var testInput = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

var _ = Describe("Day 7", func() {
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
		a := NewSolution(1)
		a.Parse(f)
		res := a.Part1()
		Expect(res).To(Equal("6440"))
	})

	It("Part 2", func() {
		err := os.WriteFile(testFileName, []byte(testInput), 0666)
		if err != nil {
			log.Fatal(err)
		}
		f := util.OpenFile(testFileName)
		defer f.Close()
		defer os.Remove(testFileName)
		a := NewSolution(2)
		a.Parse(f)
		res := a.Part2()
		Expect(res).To(Equal("5905"))
	})

})
