package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/10wpressure/advent-of-code/types"
	"github.com/10wpressure/advent-of-code/util"
)

const (
	inputFile = "input.txt"
	chunkSize = 100
)

type Race struct {
	Time types.Range
	Dist types.Range
}

func NewRace(t, d int64) *Race {
	return &Race{
		Time: types.Range{Start: 0, End: t},
		Dist: types.Range{Start: d, End: math.MaxInt64},
	}
}

func (r *Race) String() string {
	return fmt.Sprintf("Time: %+v, Dist: %+v", r.Time, r.Dist)
}

func (r *Race) Print() {
	fmt.Println(r.String())
}

func (r *Race) Calculate() int {
	counter := 0
	for i := r.Time.Start; i <= r.Time.End; i++ {
		speed := i
		time := r.Time.End - i
		dist := speed * time
		if r.Dist.ContainsNotIncluding(dist) {
			counter++
		}

	}
	return counter
}

type Solution struct {
	Races []*Race
}

func (s *Solution) Calculate() int {
	counter := 1
	for _, race := range s.Races {
		counter *= race.Calculate()
	}
	return counter
}

func NewSolution() *Solution {
	return &Solution{}
}

func (s *Solution) Parse(f *os.File) {
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	line1 := scanner.Text()
	time := strings.Fields(line1)
	scanner.Scan()
	line2 := scanner.Text()
	dist := strings.Fields(line2)
	if len(time) != len(dist) {
		_ = fmt.Errorf("time and dist not same length")
	}
	length := len(time)
	for i := 1; i < length; i++ {
		t, _ := strconv.Atoi(time[i])
		d, _ := strconv.Atoi(dist[i])
		race := NewRace(int64(t), int64(d))
		s.Races = append(s.Races, race)
	}
}

func (s *Solution) ParsePart2() {
	timeStr := ""
	distStr := ""
	for i := 0; i < len(s.Races); i++ {
		timeStr += strconv.FormatInt(s.Races[i].Time.End, 10)
		distStr += strconv.FormatInt(s.Races[i].Dist.Start, 10)
	}
	time, err := strconv.Atoi(timeStr)
	if err != nil {
		log.Fatal(err)
	}
	dist, err := strconv.Atoi(distStr)
	if err != nil {
		log.Fatal(err)
	}
	race := NewRace(int64(time), int64(dist))
	race.Print()
	s.Races = []*Race{race}
}

func (s *Solution) Part1() string {
	return strconv.FormatInt(int64(s.Calculate()), 10)
}

func (s *Solution) Part2() string {
	return strconv.FormatInt(int64(s.Calculate()), 10)
}

func main() {
	f := util.OpenFile(inputFile)
	defer f.Close()
	a := NewSolution()
	a.Parse(f)
	part1 := a.Part1()
	a.ParsePart2()
	part2 := a.Part2()

	log.Printf("Part 1: %s", part1)
	log.Printf("Part 2: %s", part2)
}
