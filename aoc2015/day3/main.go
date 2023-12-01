package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/10wpressure/advent-of-code/util"
)

const (
	chunkSize = 2
)

type Santa struct {
	Pos   Pos
	Route map[Pos]int
}

func (s *Santa) Move(v byte) Pos {
	switch v {
	case '^':
		s.Pos.y++
	case 'v':
		s.Pos.y--
	case '>':
		s.Pos.x++
	case '<':
		s.Pos.x--
	}
	s.Route[s.Pos]++
	return s.Pos
}

type Pos struct {
	x, y int64
}

func (p *Pos) String() string {
	return fmt.Sprintf("{x:%d y:%d}", p.x, p.y)

}

func part1(f *os.File) string {
	//fmt.Printf("########################################\n")
	p := Pos{0, 0}
	s := &Santa{Pos: p, Route: make(map[Pos]int)}
	s.Route[p] = 1

	util.ReadInChunks(f, chunkSize, func(buf []byte) bool {
		for _, v := range buf {
			s.Move(v)
		}
		return false
	})
	cnt := 0
	for _, v := range s.Route {
		if v > 0 {
			cnt++
			//fmt.Printf("%d: House %+v has been visited %d times\n", cnt, k, v)
		}
	}
	return strconv.Itoa(cnt)
}

func part2(f *os.File) string {
	//fmt.Printf("########################################\n")
	p := Pos{0, 0}
	s1 := &Santa{Pos: p, Route: make(map[Pos]int)}
	s2 := &Santa{Pos: p, Route: make(map[Pos]int)}
	result := &Santa{Pos: p, Route: make(map[Pos]int)}
	s1.Route[p], s2.Route[p], result.Route[p] = 1, 1, 1

	util.ReadInChunks(f, chunkSize, func(buf []byte) bool {
		for i, v := range buf {
			if i%2 == 0 {
				p = s1.Move(v)
				s1.Route[p]++
				result.Route[p]++
			} else {
				p = s2.Move(v)
				s2.Route[p]++
				result.Route[p]++
			}
		}
		return false
	})

	cnt := 0
	for _, v := range result.Route {
		if v > 0 {
			cnt++
			//fmt.Printf("S1: %d: House %+v has been visited %d times\n", cnt, k, v)
		}
	}
	return strconv.Itoa(cnt)
}

func main() {
	f := util.OpenFile("input.txt")
	defer f.Close()

	res1 := part1(f)
	log.Println("Part 1:", res1)

	f = util.OpenFile("input.txt")
	defer f.Close()

	res2 := part2(f)
	log.Println("Part 2:", res2)
}
