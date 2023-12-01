package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Suite")
}

var _ = Describe("Day 2", func() {
	It("Part 1", func() {
		box := NewBox(2, 3, 4)
		Expect(box.SurfaceArea()).To(Equal(int64(52)))
		Expect(box.SurfaceAreaWithExtra()).To(Equal(int64(58)))

		box = NewBox(1, 1, 10)
		Expect(box.SurfaceArea()).To(Equal(int64(42)))
		Expect(box.SurfaceAreaWithExtra()).To(Equal(int64(43)))
	})
	It("Part 2", func() {
		box := NewBox(2, 3, 4)
		Expect(box.CalculateRibbon()).To(Equal(int64(34)))

		box = NewBox(1, 1, 10)
		Expect(box.CalculateRibbon()).To(Equal(int64(14)))
	})
})
