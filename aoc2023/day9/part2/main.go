package part2

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
	EndsWithA  []string
	EndsWithZ  []string
	Answers    []int64
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

func EndsWithA(s string) bool {
	return strings.HasSuffix(s, "A")
}

func EndsWithZ(s string) bool {
	return strings.HasSuffix(s, "Z")
}

func (s *Solution) Parse(f *os.File) {
	scanner := bufio.NewScanner(f)
	scanner.Scan() // 1st line must be directions
	directions := scanner.Text()
	s.ParseDirections(directions)
	scanner.Scan() // 2nd line is blank

	for scanner.Scan() {
		nodeToParse := scanner.Text()
		node := s.ParseNode(nodeToParse)
		if EndsWithA(node.Name) {
			s.EndsWithA = append(s.EndsWithA, node.Name)
		}
		if EndsWithZ(node.Name) {
			s.EndsWithZ = append(s.EndsWithZ, node.Name)
		}
	}
}

func (s *Solution) Traverse(root *Node) int64 {
	var (
		i           = 0
		steps int64 = 0
		dir         = s.Directions[i]
	)

	for {
		switch dir {
		case L:
			root = s.Network[root.L]
		default:
			root = s.Network[root.R]
		}
		//fmt.Printf("%s %s\n", root.Name, dir)
		steps++
		if EndsWithZ(root.Name) {
			break
		}

		i++
		dir = s.Directions[i%len(s.Directions)]
	}
	return steps
}

func (s *Solution) Part2() string {
	for _, v := range s.EndsWithA {
		root := s.Network[v]
		answer := s.Traverse(root)
		s.Answers = append(s.Answers, answer)
	}
	res := util.LCM(s.Answers)

	return strconv.FormatInt(res, 10)
}

func Solve() string {
	f := util.OpenFile(inputFile)
	defer f.Close()
	a := NewSolution()
	a.Parse(f)
	part2 := a.Part2()
	return part2
}
