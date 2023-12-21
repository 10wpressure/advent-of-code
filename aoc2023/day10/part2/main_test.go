package part2

//import (
//	"log"
//	"os"
//	"testing"
//
//	"github.com/10wpressure/advent-of-code/util"
//	. "github.com/onsi/ginkgo"
//	. "github.com/onsi/gomega"
//)
//
//const (
//	testFileName = "test.txt"
//)
//
//func Test(t *testing.T) {
//	RegisterFailHandler(Fail)
//	RunSpecs(t, SolutionName())
//}
//
//var testInput = `0 3 6 9 12 15
//1 3 6 10 15 21
//10 13 16 21 30 45`
//
//var _ = Describe("Day 9", func() {
//	BeforeEach(func() {
//		_ = os.Remove(testFileName)
//	})
//
//	It("Part 2", func() {
//		err := os.WriteFile(testFileName, []byte(testInput), 0666)
//		if err != nil {
//			log.Fatal(err)
//		}
//		f := util.OpenFile(testFileName)
//		defer f.Close()
//		defer os.Remove(testFileName)
//		a := NewSolution()
//		a.Parse(f)
//		res := a.Part2()
//		Expect(res).To(Equal("2"))
//	})
//
//})
