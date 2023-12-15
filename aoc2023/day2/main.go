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

func SolutionName() string {
	return "Day 2: Cube Conundrum"
}

type Color string

const (
	red   Color = "red"
	green       = "green"
	blue        = "blue"
)

func (c Color) String() string {
	return string(c)
}

var Games = make([]Game, 0)

type Game struct {
	ID int
	S  []map[Color]int
}

var DefaultGame = Game{
	S: make([]map[Color]int, 0),
}

func process(f *os.File) {
	Games = make([]Game, 0)
	DefaultGame.S = append(DefaultGame.S, map[Color]int{
		red:   12,
		green: 13,
		blue:  14,
	})

	scanner := bufio.NewScanner(f)

	var iter = 0
	for scanner.Scan() {
		iter++
		line := scanner.Text()

		s := strings.Split(line, ":")

		var game Game
		g := strings.Fields(s[0])
		if g[0] == "Game" {
			num, _ := strconv.Atoi(g[1])
			game = Game{
				ID: num,
				S:  make([]map[Color]int, 0),
			}
		}

		sets := strings.Split(s[1], ";")
		for _, set := range sets {
			boxes := strings.Split(set, ",")
			hash := make(map[Color]int)
			for _, box := range boxes {
				b := strings.Fields(box)
				amount, err := strconv.Atoi(b[0])
				if err != nil {
					fmt.Printf("b0: %s\n", b[0])
					amount, err = strconv.Atoi(b[1])
					if err != nil {
						fmt.Printf("b1: %s\n", b[1])
						//fmt.Printf("Error: %+v\n", err)
					}
				}
				color := Color(b[1])
				hash[color] = amount
			}
			game.S = append(game.S, hash)
			fmt.Printf(" %+v\n", game)
		}
		Games = append(Games, game)
	}
}
func part1() string {
	result := 0

gameLoop:
	for _, curGame := range Games {
		for _, curSet := range curGame.S {
			if curSet[red] > DefaultGame.S[0][red] || curSet[green] > DefaultGame.S[0][green] || curSet[blue] > DefaultGame.S[0][blue] {
				continue gameLoop
			}
		}
		result += curGame.ID
	}

	return strconv.Itoa(result)
}

func part2() string {
	result := 0

	for _, curGame := range Games {
		minSet := map[Color]int{
			red:   curGame.S[0][red],
			green: curGame.S[0][green],
			blue:  curGame.S[0][blue],
		}
		for _, curSet := range curGame.S {
			if curSet[red] > minSet[red] {
				minSet[red] = curSet[red]
			}
			if curSet[green] > minSet[green] {
				minSet[green] = curSet[green]
			}
			if curSet[blue] > minSet[blue] {
				minSet[blue] = curSet[blue]
			}
		}
		power := minSet[red] * minSet[green] * minSet[blue]
		result += power
	}

	return strconv.Itoa(result)
}

func main() {
	f := util.OpenFile("input.txt")
	defer f.Close()
	process(f)
	result := part1()
	fmt.Println("Part 1:", result)

	result = part2()
	fmt.Println("Part 2:", result)

}
