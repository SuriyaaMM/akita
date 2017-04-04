package core_test

import (
	"gitlab.com/yaotsu/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Freq", func() {
	It("should get period", func() {
		var f core.Freq = 1 * core.GHz
		Expect(f.Period()).To(BeNumerically("==", 1e-9))
	})

	It("should get the next tick", func() {
		var f core.Freq = 1 * core.GHz
		Expect(f.NextTick(102.000000001)).To(BeNumerically("~", 102.000000002, 1e-12))
	})

	It("should get the next tick 2", func() {
		var f core.Freq = 1 * core.GHz
		Expect(f.NextTick(0.000000031)).To(BeNumerically("~", 0.000000032, 1e-12))
	})

	It("should get the next tick 2", func() {
		var f core.Freq = 1 * core.GHz
		Expect(f.NextTick(0.000000017)).To(BeNumerically("~", 0.000000018, 1e-12))
	})

	It("should get the next tick, if currTime is not on a tick", func() {
		var f core.Freq = 1 * core.GHz
		Expect(f.NextTick(102.0000000011)).To(BeNumerically("~", 102.000000002, 1e-12))
	})

	It("should get the n cycles later", func() {
		var f core.Freq = 1 * core.GHz
		Expect(f.NCyclesLater(12, 102.000000001)).To(
			BeNumerically("~", 102.000000014, 1e-12))
	})

	It("should get the n cycles later, if current time is not on a tick", func() {
		var f core.Freq = 1 * core.GHz
		Expect(f.NCyclesLater(12, 102.0000000011)).To(
			BeNumerically("~", 102.000000014, 1e-12))
	})

	It("should get the no-earlier-than time, on tick", func() {
		var f core.Freq = 1 * core.GHz
		Expect(f.NoEarlierThan(102.00)).To(BeNumerically("~", 102.00, 1e-12))
	})

	It("should get the no-earlier-than time, off tick", func() {
		var f core.Freq = 1 * core.GHz
		Expect(f.NoEarlierThan(102.0000000011)).To(BeNumerically("~", 102.000000002, 1e-12))
	})

})
