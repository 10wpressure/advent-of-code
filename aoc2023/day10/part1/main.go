package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	t "github.com/10wpressure/advent-of-code/aoc2023/day10/types"
	"github.com/10wpressure/advent-of-code/util"
)

const (
	inputFile = "input.txt"
	chunkSize = 100
)

type Solution struct {
	Root    *t.Pipe
	Pipes   map[t.Point]*t.Pipe
	Dirs    map[t.Direction]bool
	Visited map[t.Point]bool
}

func SolutionName() string {
	return "Day 10: Pipe Maze"
}

func NewSolution() *Solution {
	return &Solution{
		Root:    nil,
		Pipes:   make(map[t.Point]*t.Pipe),
		Dirs:    make(map[t.Direction]bool),
		Visited: make(map[t.Point]bool),
	}
}

func (s *Solution) SetRootTile() {
	s.Dirs[t.South] = s.Pipes[t.Point{X: s.Root.Point.X, Y: s.Root.Point.Y + 1}].Allowed(t.North)
	s.Dirs[t.North] = false
	s.Dirs[t.East] = s.Pipes[t.Point{X: s.Root.Point.X + 1, Y: s.Root.Point.Y}].Allowed(t.West)
	s.Dirs[t.West] = false

	if s.Root.Point.X > 0 {
		s.Dirs[t.West] = s.Pipes[t.Point{X: s.Root.Point.X - 1, Y: s.Root.Point.Y}].Allowed(t.East)
	}
	if s.Root.Point.Y > 0 {
		s.Dirs[t.North] = s.Pipes[t.Point{X: s.Root.Point.X, Y: s.Root.Point.Y - 1}].Allowed(t.South)
	}

	if s.Dirs[t.South] && s.Dirs[t.North] && !s.Dirs[t.East] && !s.Dirs[t.West] {
		s.Root.Tile = t.NorthSouth
	} else if s.Dirs[t.East] && s.Dirs[t.West] && !s.Dirs[t.South] && !s.Dirs[t.North] {
		s.Root.Tile = t.EastWest
	} else if s.Dirs[t.South] && s.Dirs[t.East] && !s.Dirs[t.North] && !s.Dirs[t.West] {
		s.Root.Tile = t.SouthEast
	} else if s.Dirs[t.North] && s.Dirs[t.West] && !s.Dirs[t.South] && !s.Dirs[t.East] {
		s.Root.Tile = t.NorthWest
	} else if s.Dirs[t.North] && s.Dirs[t.East] && !s.Dirs[t.South] && !s.Dirs[t.West] {
		s.Root.Tile = t.NorthEast
	} else if s.Dirs[t.South] && s.Dirs[t.West] && !s.Dirs[t.North] && !s.Dirs[t.East] {
		s.Root.Tile = t.SouthWest
	} else {
		log.Fatalln("Could not find root tile, to many directions: ", s.Dirs)
	}
}

func (s *Solution) Parse(f *os.File) {
	scanner := bufio.NewScanner(f)
	var y = 0
	for scanner.Scan() {
		parsed := strings.Split(scanner.Text(), "")
		for x, char := range parsed {
			p := t.Point{
				X: int64(x),
				Y: int64(y),
			}
			pipe := t.NewPipe(char, p)
			s.Pipes[p] = &pipe
			if pipe.Tile == t.Start {
				s.Root = &pipe
				fmt.Printf("ROOT: %+v\n", s.Root)
			}
		}
		y++
	}
	s.SetRootTile()
}

func (s *Solution) Part1() string {
	curPos := s.Root
	var alreadyVisited = false
	var curDir = t.West
	for i := 0; i < len(s.Dirs); i++ {
		if s.Dirs[t.Direction(i)] {
			curDir = t.Direction(i)
			break
		}
	}

	for !alreadyVisited {
		s.Visited[curPos.Point] = true
		delta := curDir.Delta()
		curPos = s.Pipes[t.Point{X: curPos.Point.X + delta.X, Y: curPos.Point.Y + delta.Y}]
		curDir = curPos.Next(curDir)
		fmt.Printf("Pos: %+v, Dir: %+v, Delta: %+v\n", curPos, curDir, delta)
		alreadyVisited = s.Visited[curPos.Point]
	}

	return strconv.Itoa(len(s.Visited) / 2)
}

func Solve() string {
	f := util.OpenFile(inputFile)
	defer f.Close()
	a := NewSolution()
	a.Parse(f)
	part1 := a.Part1()
	return part1
}
