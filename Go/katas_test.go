package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "katas"
)

var _ = Describe("Tests using hard-coded strings", func() {
	It("Should return a correctly translated string", func() {
		Expect(ToNato("If you can read")).To(Equal("India Foxtrot Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta"))
		Expect(ToNato("Did not see that coming")).To(Equal("Delta India Delta November Oscar Tango Sierra Echo Echo Tango Hotel Alfa Tango Charlie Oscar Mike India November Golf"))
		Expect(ToNato("go for it!")).To(Equal("Golf Oscar Foxtrot Oscar Romeo India Tango !"))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1})).To(Equal([]int{1, 2, 1, 1, 3, 1, 0, 0, 0, 0}))
	})
})

func dotest(code, msg, expected string) {
	Expect(HamsterMe(code, msg)).To(Equal(expected), "With code = \"%s\", message = \"%s\"", code, msg)
}

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		dotest("hamster", "hamster", "h1a1m1s1t1e1r1")
		dotest("hamster", "helpme", "h1e1h5m4m1e1")
		dotest("hmster", "hamster", "h1t8m1s1t1e1r1")
		dotest("hhhhammmstteree", "hamster", "h1a1m1s1t1e1r1")
		dotest("hamster", "abcdefghijklmnopqrstuvwxyz", "a1a2a3a4e1e2e3h1h2h3h4h5m1m2m3m4m5r1s1t1t2t3t4t5t6t7")
		dotest("f", "abcdefghijklmnopqrstuvwxyz", "f22f23f24f25f26f1f2f3f4f5f6f7f8f9f10f11f12f13f14f15f16f17f18f19f20f21")
	})
})

func dotest2(n int, exp int) {
	var ans = DblLinear(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest2(50, 175)
		dotest2(100, 447)
		dotest2(500, 3355)
		dotest2(1000, 8488)
	})
})
