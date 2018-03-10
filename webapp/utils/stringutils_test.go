package utils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/gesiel/go-collect/webapp/utils"
)

var _ = Describe("String utils", func() {
	Context("Should validate empty values", func() {
		It("Empty with no spaces", func() {
			Expect(IsValidValue("")).Should(BeFalse())
		})
		It("Empty with spaces", func() {
			Expect(IsValidValue("   ")).Should(BeFalse())
		})
		It("Valid value with no spaces", func() {
			Expect(IsValidValue("valid")).Should(BeTrue())
		})
		It("Valid value with spaces", func() {
			Expect(IsValidValue("valid")).Should(BeTrue())
		})
	})
})
