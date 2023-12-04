package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/10wpressure/advent-of-code/util"
)

const (
	fileName  = "input.txt"
	chunkSize = 100
)

type Game struct {
	ID      int
	Winning Set
	Owned   []int
	Score   int
	Doubled bool
}

type Set map[int]struct{}

// Add Adds a key to the set
func (s Set) Add(num int) {
	s[num] = struct{}{}
}

// Remove Removes a key from the set
func (s Set) Remove(num int) {
	delete(s, num)
}

// Clear Removes all keys from the set
func (s Set) Clear() {
	for k := range s {
		delete(s, k)
	}
}

// Len Returns the number of keys in the set
func (s Set) Len() int {
	return len(s)
}

// Has Returns a boolean value describing if the key exists in the set
func (s Set) Has(num int) bool {
	_, ok := s[num]
	return ok
}

func process(f *os.File) []Game {
	games := make([]Game, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		winningSet := make(Set)
		ownedSlice := make([]int, 0)

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
			ownedSlice = append(ownedSlice, digit)
		}

		games = append(games, Game{
			ID:      id,
			Winning: winningSet,
			Owned:   ownedSlice,
			Score:   0,
			Doubled: false,
		})
	}
	return games
}

func part1(f *os.File) string {
	result := 0
	games := process(f)
	for k := 0; k < len(games); k++ {
		for j := 0; j < len(games[k].Owned); j++ {
			if games[k].Winning.Has(games[k].Owned[j]) {
				if !games[k].Doubled {
					games[k].Score = 1
					games[k].Doubled = true
				} else {
					fmt.Printf("Game %d has already been doubled\n", games[k].ID)
					games[k].Score = games[k].Score * 2
				}
			}
		}
	}

	for _, v := range games {
		result += v.Score
	}
	return strconv.Itoa(result)
}

func part2(f *os.File) string {
	return ""
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
