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

var _ = Describe("Day 1", func() {
	BeforeEach(func() {
		_ = os.Remove(testFileName)
	})

	It("Part 1 Test 1", func() {
		err := os.WriteFile(testFileName, []byte("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"), 0666)
		if err != nil {
			log.Fatal(err)
		}
		f := util.OpenFile(testFileName)
		defer f.Close()
		defer os.Remove(testFileName)
		res := part1(f)
		Expect(res).To(Equal("13"))
	})

	It("Part 2 Test 1", func() {
		//err := os.WriteFile(testFileName, []byte("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"), 0666)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//f := util.OpenFile(testFileName)
		//defer f.Close()
		//defer os.Remove(testFileName)
		//res := part2(f)
		//s := 467835
		//Expect(res).To(Equal(strconv.Itoa(s)))
	})

})
