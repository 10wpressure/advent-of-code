package main

import (
	"bufio"
	"fmt"
	"math"
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
	return "Day 1: Trebuchet?!"
}

var Digits = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var LiteralDigits = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func FindFirstAndLastDigit(s string) int {
	var first = make(map[int]int)
	var last = make(map[int]int)
	fd, ld := 0, 0
	for k, v := range Digits {
		ind := strings.Index(s, k)
		indL := strings.LastIndex(s, k)
		if ind == -1 {
			continue
		}
		first[ind] = v
		last[indL] = v
	}
	for k, v := range LiteralDigits {
		ind := strings.Index(s, k)
		indL := strings.LastIndex(s, k)
		if ind == -1 {
			continue
		}
		first[ind] = v
		last[indL] = v
	}

	minFirstKey := math.MaxInt64
	for k := range first {
		if k < minFirstKey {
			minFirstKey = k
		}
	}

	maxLastKey := math.MinInt64
	for k := range last {
		if k > maxLastKey {
			maxLastKey = k
		}
	}

	fd = first[minFirstKey]
	ld = last[maxLastKey]

	return fd*10 + ld
}

func part1(f *os.File) string {
	scanner := bufio.NewScanner(f)
	var res = 0

	var iter = 0
	for scanner.Scan() {
		iter++
		line := []rune(scanner.Text())
		fd := 0
		ld := 0
		for _, v := range line {
			n, err := strconv.Atoi(string(v))
			if err != nil {
				continue
			}
			//_ = fmt.Sprintf("%d: %c\n", i, v)
			if fd == 0 {
				fd = n * 10
				ld = n
				continue
			} else {
				ld = n
			}
		}
		res += fd + ld
	}
	return strconv.Itoa(res)
}

func part2(f *os.File) string {
	scanner := bufio.NewScanner(f)
	var res = 0

	var iter = 0
	for scanner.Scan() {
		line := scanner.Text()
		iter++
		cur := FindFirstAndLastDigit(line)
		res += cur
	}
	return strconv.Itoa(res)
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
