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
		err := os.WriteFile(testFileName, []byte("467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."), 0666)
		if err != nil {
			log.Fatal(err)
		}
		f := util.OpenFile(testFileName)
		defer f.Close()
		defer os.Remove(testFileName)
		res := part1(f)
		Expect(res).To(Equal("4361"))
	})

	It("Part 2 Test 1", func() {
		err := os.WriteFile(testFileName, []byte("467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."), 0666)
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
