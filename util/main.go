package util

import (
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/constraints"
)

func ReadInChunks(f *os.File, chunkSize int, cb func([]byte) bool) {
	buf := make([]byte, chunkSize)

	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		if n == 0 {
			break
		}

		// Pass only the relevant portion of the buffer to the callback
		br := cb(buf[:n])
		if br {
			break
		}

		if err == io.EOF {
			break
		}
	}
}

func OpenFile(path string) *os.File {
	if strings.HasPrefix(path, "../") {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		wd = strings.TrimSuffix(wd, "/part1")
		wd = strings.TrimSuffix(wd, "/part2")
		path = strings.Join([]string{wd, path}, "/")
	}

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func FileIntoString(f *os.File) string {
	var buf strings.Builder
	_, err := io.Copy(&buf, f)
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}

type Set[T comparable] map[T]struct{}

// Add Adds a key to the set
func (s Set[T]) Add(key T) {
	s[key] = struct{}{}
}

// Remove Removes a key from the set
func (s Set[T]) Remove(key T) {
	delete(s, key)
}

// Clear Removes all keys from the set
func (s Set[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

// Len Returns the number of keys in the set
func (s Set[T]) Len() int {
	return len(s)
}

// Has Returns a boolean value describing if the key exists in the set
func (s Set[T]) Has(key T) bool {
	_, ok := s[key]
	return ok
}

func intersection(s1, s2 []string) (inter []string) {
	hash := make(map[string]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		// If elements present in the hashmap then append intersection list.
		if hash[e] {
			inter = append(inter, e)
		}
	}
	//Remove dups from slice.
	inter = removeDups(inter)
	return
}

// Remove dups from slice.
func removeDups(elements []string) (nodups []string) {
	encountered := make(map[string]bool)
	for _, element := range elements {
		if !encountered[element] {
			nodups = append(nodups, element)
			encountered[element] = true
		}
	}
	return
}

type Number interface {
	constraints.Integer
}

func GCD[T Number](a, b T) T {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM[T Number](integers []T) T {
	if len(integers) < 2 {
		panic("LCM requires at least two integers")
	}

	result := integers[0]
	for i := 1; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

func lcm[T Number](a, b T) T {
	return a * b / GCD(a, b)
}
