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
	Winning util.Set[int]
	Owned   []int
	Score   int
	Doubled bool
	Matches int
}

func process(f *os.File) []Game {
	games := make([]Game, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		winningSet := make(util.Set[int])
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
			Matches: 0,
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
	result := 0
	games := process(f)
	for k := 0; k < len(games); k++ {
		for j := 0; j < len(games[k].Owned); j++ {
			if games[k].Winning.Has(games[k].Owned[j]) {
				games[k].Matches++
			}
		}
	}

	for k := 0; k < len(games); k++ {
		for m := 1; m <= games[k].Matches; m++ {
			if k+m >= len(games) {
				continue
			}
			games[k+m].Matches++
		}
	}

	for k := 0; k < len(games); k++ {
		fmt.Printf("Game %d has %d matches\n", games[k].ID, games[k].Matches)
		result += games[k].Matches
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
