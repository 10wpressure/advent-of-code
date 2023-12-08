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
	RunSpecs(t, "API Suite")
}

var testInput = "seeds: 79 14 55 13\n\nseed-to-soil map:\n50 98 2\n52 50 48\n\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\n\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\n\nwater-to-light map:\n88 18 7\n18 25 70\n\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\n\ntemperature-to-humidity map:\n0 69 1\n1 0 69\n\nhumidity-to-location map:\n60 56 37\n56 93 4"

var _ = Describe("Day 5", func() {
	BeforeEach(func() {
		_ = os.Remove(testFileName)
	})

	It("Part 1 Test 1", func() {
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

	It("Part 2 Test 1", func() {
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
