package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	// test disabled because runner module was removed, will be fixed later
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Suite")
}

var _ = Describe("Day 1", func() {
	It("Part 1", func() {
		Expect(part1()).To(Equal("232"))
	})
	It("Part 2", func() {
		Expect(part2()).To(Equal("1783"))
	})
})
