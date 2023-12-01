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
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Suite")
}

var _ = Describe("Day 1", func() {
	BeforeEach(func() {
		_ = os.Remove(testFileName)
	})

	It("Part 1 Test 1", func() {
		err := os.WriteFile(testFileName, []byte("1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"), 0666)
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
		err := os.WriteFile(testFileName, []byte("two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen\nnineeight6nine1three1eight"), 0666)
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
