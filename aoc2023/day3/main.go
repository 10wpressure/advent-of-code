package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/10wpressure/advent-of-code/util"
)

const (
	fileName  = "input.txt"
	chunkSize = 100
)

type SymbolSet map[string]struct{}

// Add Adds a key to the set
func (s SymbolSet) Add(symbol string) {
	s[symbol] = struct{}{}
}

// Remove Removes a key from the set
func (s SymbolSet) Remove(symbol string) {
	delete(s, symbol)
}

// Clear Removes all keys from the set
func (s SymbolSet) Clear() {
	for k := range s {
		delete(s, k)
	}
}

// Len Returns the number of keys in the set
func (s SymbolSet) Len() int {
	return len(s)
}

// Has Returns a boolean value describing if the key exists in the set
func (s SymbolSet) Has(symbol string) bool {
	_, ok := s[symbol]
	return ok
}

type Coord struct {
	x, y int
}

type Number struct {
	Coord
	Value int
}

func (n Number) String() string {
	return strconv.Itoa(n.Value)
}

func (n Number) Len() int {
	return len(n.String())
}

type Symbol struct {
	Coord
	Value   string
	Details []Number
}

type Space struct {
	Coord
}

func process(f *os.File) (numbers []Number, symbols map[Coord]Symbol, spaces []Space) {
	symbols = make(map[Coord]Symbol)
	scanner := bufio.NewScanner(f)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		acc := ""
		ind := -1
		// iterate over the line
		for x, s := range line {
			// check if the char is a number
			number, err := strconv.Atoi(string(s))
			if err == nil {
				if acc == "" {
					ind = x
				}
				acc += string(s)
			} else {
				if acc != "" {
					num, err := strconv.Atoi(acc)
					if err != nil {
						log.Printf("acc must be a number: %s", acc)
					}
					number = num

					numbers = append(numbers, Number{
						Coord: Coord{
							x: ind,
							y: y,
						},
						Value: number,
					})
					ind = -1
					acc = ""
				}

				if string(s) == "." {
					spaces = append(spaces, Space{
						Coord: Coord{
							x: x,
							y: y,
						},
					})
					continue
				}

				symbol := Symbol{
					Coord: Coord{
						x: x,
						y: y,
					},
					Value:   string(s),
					Details: make([]Number, 0),
				}
				symbols[symbol.Coord] = symbol
			}
		}
		y++
	}
	return numbers, symbols, spaces
}

func part1(f *os.File) string {
	numbers, symbols, _ := process(f)

	result := 0
	for i, n := range numbers {
		fmt.Printf("number %d: %+v %+v\n", i, n.Coord, n.Value)

		for y := n.Coord.y - 1; y <= n.Coord.y+1; y++ {
			for x := n.Coord.x - 1; x <= n.Coord.x+n.Len(); x++ {
				if y == n.Coord.y && x == n.Coord.x {
					continue
				}
				_, ok := symbols[Coord{
					x: x,
					y: y,
				}]
				if !ok {
					continue
				} else {
					result += n.Value
				}
			}
		}
	}

	return strconv.Itoa(result)
}

func part2(f *os.File) string {
	numbers, symbols, _ := process(f)

	result := 0
	for i, n := range numbers {
		fmt.Printf("number %d: %+v %+v\n", i, n.Coord, n.Value)

		for y := n.Coord.y - 1; y <= n.Coord.y+1; y++ {
			for x := n.Coord.x - 1; x <= n.Coord.x+n.Len(); x++ {
				if y == n.Coord.y && x == n.Coord.x {
					continue
				}
				symbol, ok := symbols[Coord{
					x: x,
					y: y,
				}]
				if !ok {
					continue
				}

				if symbol.Value == "*" {
					symbol.Details = append(symbol.Details, n)
					symbols[Coord{x: x, y: y}] = symbol
					if len(symbol.Details) == 2 {
						gearRatio := symbol.Details[0].Value * symbol.Details[1].Value
						result += gearRatio
					}
					fmt.Printf("symbol: %+v\n", symbol)
					fmt.Println("Details: ", symbol.Details)
				}
			}
		}
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
