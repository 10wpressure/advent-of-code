package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/10wpressure/advent-of-code/util"
	"github.com/emirpasic/gods/sets/linkedhashset"
)

const (
	fileName  = "input.txt"
	chunkSize = 100
)

type Card struct {
	ID      int
	Winning *linkedhashset.Set
	Owned   *linkedhashset.Set
	Score   int
	Doubled bool
	Matches int
}

func (c *Card) Count() int {
	return c.Winning.Intersection(c.Owned).Size()
}

func (c *Card) CalculateScore() int {
	for j := 0; j < c.Owned.Size(); j++ {
		if c.Winning.Contains(c.Owned.Values()[j]) {
			if !c.Doubled {
				c.Score = 1
				c.Doubled = true
			} else {
				c.Score = c.Score * 2
			}
		}
	}
	return c.Score
}

func process(f *os.File) []*Card {
	games := make([]*Card, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		winningSet := linkedhashset.New()
		ownedSet := linkedhashset.New()

		line := scanner.Text()
		s := strings.Split(line, ":")
		id, err := strconv.Atoi(strings.Fields(s[0])[1])
		if err != nil {
			_ = fmt.Errorf("failed to parse id: %s", err)
		}
		halves := strings.Split(s[1], "|")
		winning := strings.Fields(halves[0])
		owned := strings.Fields(halves[1])

		for _, symbol := range winning {
			digit, err := strconv.Atoi(symbol)
			if err != nil {
				_ = fmt.Errorf("failed to parse digit: %s", err)
			}
			winningSet.Add(digit)
		}
		for _, symbol := range owned {
			digit, err := strconv.Atoi(symbol)
			if err != nil {
				_ = fmt.Errorf("failed to parse digit: %s", err)
			}
			ownedSet.Add(digit)
		}

		games = append(games, &Card{
			ID:      id,
			Winning: winningSet,
			Owned:   ownedSet,
			Score:   0,
			Doubled: false,
			Matches: 0,
		})
	}
	return games
}

func part1(f *os.File) string {
	result := 0
	cards := process(f)

	for k := 0; k < len(cards); k++ {
		cards[k].CalculateScore()
	}

	for _, v := range cards {
		result += v.Score
	}

	return strconv.Itoa(result)
}

func part2(f *os.File) string {
	result := 0
	cards := process(f)

	var multiplier = make([]int, 0)
	for i := 0; i < len(cards); i++ {
		multiplier = append(multiplier, 1)
	}

	for i, card := range cards {
		count := card.Count()
		for j := i + 1; j < i+1+count; j++ {
			multiplier[j] += multiplier[i]
		}
	}

	for i := range multiplier {
		result += multiplier[i]
	}

	return strconv.Itoa(result)
}

func main() {
	f := util.OpenFile("input.txt")
	defer f.Close()
	result := part1(f)
	fmt.Println("Part 1:", result)

	f = util.OpenFile(fileName)
	defer f.Close()
	result = part2(f)
	fmt.Println("Part 2:", result)
}
