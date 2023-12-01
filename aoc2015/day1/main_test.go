package main

import (
	"testing"

	"github.com/10wpressure/advent-of-code/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Suite")
}

var _ = Describe("Day 1", func() {
	f := util.OpenFile("input.txt")
	defer f.Close()

	It("Part 1", func() {
		Expect(part1(f)).To(Equal("232"))
	})
	It("Part 2", func() {
		Expect(part2(f)).To(Equal("1783"))
	})
})
