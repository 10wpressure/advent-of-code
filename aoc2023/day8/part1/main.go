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

type Direction int

const (
	None Direction = iota
	L
	R
)

func (d Direction) String() string {
	switch d {
	case L:
		return "L"
	case R:
		return "R"
	default:
		log.Fatalf("invalid direction: %d", d)
		return ""
	}
}

func FromString(r rune) Direction {
	switch r {
	case 'L':
		return L
	case 'R':
		return R
	default:
		log.Fatalf("invalid direction: %c", r)
		return None
	}
}

type Solution struct {
	Root       *Node
	Directions []Direction
	Network    map[string]*Node
	Steps      int64
}

type Node struct {
	Name string
	L    string
	R    string
}

func SolutionName() string {
	return "Day 8: Haunted Wasteland"
}

func NewSolution() *Solution {
	return &Solution{Directions: make([]Direction, 0), Network: make(map[string]*Node)}
}

func (s *Solution) ParseDirections(input string) {
	for _, v := range input {
		s.Directions = append(s.Directions, FromString(v))
	}
	//fmt.Printf("%v\n", s.Directions)
}

func (s *Solution) ParseNode(input string) *Node {
	nodeStr := strings.Fields(input)
	key := nodeStr[0]
	left := strings.Trim(nodeStr[2], "(,)")
	right := strings.Trim(nodeStr[3], "(,)")
	//fmt.Printf("key: %s, left: %s, right: %s\n", key, left, right)
	s.Network[key] = &Node{
		Name: key,
		L:    left,
		R:    right,
	}

	return s.Network[key]
}

func (s *Solution) Parse(f *os.File) {
	scanner := bufio.NewScanner(f)
	scanner.Scan() // 1st line must be directions
	directions := scanner.Text()
	s.ParseDirections(directions)
	scanner.Scan() // 2nd line is blank

	for scanner.Scan() {
		nodeToParse := scanner.Text()
		s.ParseNode(nodeToParse)
	}
	s.Root = s.Network["AAA"]
}

func (s *Solution) Traverse() int64 {
	var i = 0
	var dir = s.Directions[i]

	for {
		root := s.Network[s.Root.Name]
		switch dir {
		case L:
			s.Root = s.Network[root.L]
		default:
			s.Root = s.Network[root.R]
		}
		//fmt.Printf("%s %s\n", s.Root.Name, dir)
		s.Steps++
		if s.Root.Name == "ZZZ" {
			break
		}

		i++
		dir = s.Directions[i%len(s.Directions)]
	}
	return s.Steps
}

func (s *Solution) Part1() string {
	res := s.Traverse()
	return strconv.FormatInt(res, 10)
}

func (s *Solution) Part2() string {
	res := s.Traverse()
	return strconv.FormatInt(res, 10)
}

func Solve() string {
	f := util.OpenFile(inputFile)
	defer f.Close()
	a := NewSolution()
	a.Parse(f)
	part1 := a.Part1()
	return part1
}
