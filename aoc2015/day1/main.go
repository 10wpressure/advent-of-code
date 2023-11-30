package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	chunkSize = 100
)

func ReadInChunks(f *os.File, chunkSize int, cb func([]byte) bool) {
	buf := make([]byte, chunkSize)
	br := false

	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		if err == io.EOF {
			break
		}
		if n == 0 {
			break
		}

		br = cb(buf)
		if br {
			break
		}
	}
}

func part1() string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	floor := 0

	ReadInChunks(f, chunkSize, func(buf []byte) bool {
		for i := range buf {
			switch string(buf[i]) {
			case "(":
				floor++
			case ")":
				floor--
			}
		}
		return false
	})

	return fmt.Sprintf("%d", floor)
}

func part2() string {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	floor := 0
	pos := 0

	ReadInChunks(f, chunkSize, func(buf []byte) bool {
		for i := 0; i < len(buf); i++ {
			switch string(buf[i]) {
			case "(":
				floor++
			case ")":
				floor--
			}
			pos++
			if floor < 0 {
				return true
			}
		}
		return false
	})

	return fmt.Sprintf("%d", pos)
}

func main() {
	res1 := part1()
	log.Println("Part 1:", res1)

	res2 := part2()
	log.Println("Part 2:", res2)

}
