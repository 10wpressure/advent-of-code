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

var testInput = `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`

var _ = Describe(SolutionName(), func() {
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
		Expect(res).To(Equal("8"))
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
		_ = a.Part1()
		res := a.Part2()
		Expect(res).To(Equal("1"))
	})
})
