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

var testInput = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`

var _ = Describe("Day 8", func() {
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
		Expect(res).To(Equal("6"))
	})

	//It("Part 2", func() {
	//	err := os.WriteFile(testFileName, []byte(testInput), 0666)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	f := util.OpenFile(testFileName)
	//	defer f.Close()
	//	defer os.Remove(testFileName)
	//	a := NewSolution(2)
	//	a.Parse(f)
	//	res := a.Part2()
	//	Expect(res).To(Equal("5905"))
	//})

})
