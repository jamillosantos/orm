package document

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GoName", func() {
	Describe("GoNamePublic", func() {
		It("should get a name from a single word", func() {
			Expect(GoNamePublic("name")).To(Equal("Name"))
		})

		It("should get a name from a many words", func() {
			Expect(GoNamePublic("this is the name")).To(Equal("ThisIsTheName"))
		})

		It("should get a name from a snake case", func() {
			Expect(GoNamePublic("this_is_the_name")).To(Equal("ThisIsTheName"))
		})
	})

	Describe("GoNamePrivate", func() {
		It("should get a name from a single word", func() {
			Expect(GoNamePrivate("name")).To(Equal("name"))
		})

		It("should get a name from a many words", func() {
			Expect(GoNamePrivate("this is the name")).To(Equal("thisIsTheName"))
		})

		It("should get a name from a snake case", func() {
			Expect(GoNamePrivate("this_is_the_name")).To(Equal("thisIsTheName"))
		})
	})
})
