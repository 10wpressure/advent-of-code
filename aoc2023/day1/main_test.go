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
	testInput1   = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	testInput2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
nineeight6nine1three1eight`
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, SolutionName())
}

var _ = Describe("Day 1", func() {
	BeforeEach(func() {
		_ = os.Remove(testFileName)
	})

	It("Part 1 Test 1", func() {
		err := os.WriteFile(testFileName, []byte(testInput1), 0666)
		if err != nil {
			log.Fatal(err)
		}
		f := util.OpenFile(testFileName)
		defer f.Close()
		defer os.Remove(testFileName)
		res := part1(f)
		Expect(res).To(Equal("142"))
	})

	It("Part 2 Test 1", func() {
		err := os.WriteFile(testFileName, []byte(testInput2), 0666)
		if err != nil {
			log.Fatal(err)
		}
		f := util.OpenFile(testFileName)
		defer f.Close()
		defer os.Remove(testFileName)
		res := part2(f)
		s := 281 + 98
		Expect(res).To(Equal(strconv.Itoa(s)))
	})

})
