package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/10wpressure/advent-of-code/util"
)

const (
	inputFile = "input.txt"
	chunkSize = 100
)

type Solution struct {
	hands []Hand
	part  int
}

func SolutionName() string {
	return "Day 7: Camel Cards"
}

func NewSolution(part int) *Solution {
	return &Solution{part: part}
}

func (s *Solution) Parse(f *os.File) {
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		cardsStr, bidStr := strings.Fields(scanner.Text())[0], strings.Fields(scanner.Text())[1]
		bid, err := strconv.ParseInt(bidStr, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		var cards = make([]Card, 0)
		for _, v := range cardsStr {
			var card Card
			if s.part == 2 {
				card = FromStringPart2(v)
			} else {
				card = FromString(v)
			}
			cards = append(cards, card)
		}
		hand := NewHand(cards, bid)
		s.hands = append(s.hands, *hand)
	}
}

func (s *Solution) Len() int      { return len(s.hands) }
func (s *Solution) Swap(i, j int) { s.hands[i], s.hands[j] = s.hands[j], s.hands[i] }
func (s *Solution) Less(i, j int) bool {
	a := &s.hands[i]
	b := &s.hands[j]

	if a.handType != b.handType {
		return a.handType.Less(b.handType)
	}

	for k := 0; k < 5; k++ {
		if a.cards[k] != b.cards[k] {
			return a.cards[k] < b.cards[k]
		}
	}
	return false
}

func (s *Solution) Part1() string {
	sort.Sort(s)
	res := s.Calculate()
	return strconv.FormatInt(res, 10)
}

func (s *Solution) Part2() string {
	sort.Sort(s)
	res := s.Calculate()
	return strconv.FormatInt(res, 10)
}
func (s *Solution) Calculate() int64 {
	var acc int64 = 0
	for i, v := range s.hands {
		s.hands[i].score = int64(i+1) * v.bid
		acc += s.hands[i].score
		//fmt.Printf("%v | %v: %v * %v = %v\n", acc, v.String(), v.bid, int64(i+1), s.hands[i].score)
	}
	return acc
}

func main() {
	f := util.OpenFile(inputFile)
	defer f.Close()
	a := NewSolution(1)
	a.Parse(f)
	part1 := a.Part1()
	f = util.OpenFile(inputFile)
	defer f.Close()
	b := NewSolution(2)
	b.Parse(f)
	part2 := b.Part2()

	log.Printf("Part 1: %s", part1)
	log.Printf("Part 2: %s", part2)
}
