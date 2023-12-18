package part1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/10wpressure/advent-of-code/util"
)

const (
	inputFile = "input.txt"
	chunkSize = 100
)

type Solution struct {
	Histories [][]int64
}

func SolutionName() string {
	return "Day 9: Mirage Maintenance"
}

func NewSolution() *Solution {
	return &Solution{}
}

func (s *Solution) Parse(f *os.File) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parsed := strings.Fields(scanner.Text())
		history := make([][]int64, 1)
		history[0] = make([]int64, len(parsed))
		for i, v := range parsed {
			n, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			history[0][i] = n
			//fmt.Printf("%v: %v\n", i, history[0][i])
		}
		s.Histories = append(s.Histories, history...)
	}
}

func (s *Solution) Part1() string {
	var predictions = make([]int64, 0)
	for _, h := range s.Histories {
		diffs := GetAllDiffsForOneHistory(h)

		diffs[len(diffs)-1] = append(diffs[len(diffs)-1], 0)
		for i := len(diffs) - 2; i >= 0; i-- {
			additionalValue := diffs[i][len(diffs[i])-1] + diffs[i+1][len(diffs[i+1])-1]
			diffs[i] = append(diffs[i], additionalValue)
		}
		predictions = append(predictions, diffs[0][len(diffs[0])-1])
	}

	return strconv.FormatInt(util.SumSlice(predictions), 10)
}

func GetAllDiffsForOneHistory(h []int64) [][]int64 {
	diffs := make([][]int64, 0)
	diffs = append(diffs, h)

	for {
		cur := diffs[len(diffs)-1]
		d := GetDiffsForOneSlice(cur)
		diffs = append(diffs, d)

		if util.SliceContainsOnlyZeroes(d) {
			break
		}
	}
	return diffs
}

func GetDiffsForOneSlice(nums []int64) []int64 {
	var diffs []int64
	for i := 1; i < len(nums); i++ {
		diffs = append(diffs, nums[i]-nums[i-1])
	}

	return diffs
}

func Solve() string {
	f := util.OpenFile(inputFile)
	defer f.Close()
	a := NewSolution()
	a.Parse(f)
	part1 := a.Part1()
	return part1
}
