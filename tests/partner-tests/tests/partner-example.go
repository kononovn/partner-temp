package tests

import (
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Partner test cases", func() {

	BeforeEach(func() {
		By("Partners BeforeEach")
	})

	Context("Partner context 1,", func() {

		It("First partner test case", func() {
			By("Partner's Test Case")

		})
	})

	Context("Partner context 2,", func() {

		It("Second Partner test case", func() {
			By("Partner's Test Case")
		})
	})
})
