package main

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Error Is", func() {
	Describe("Error Is", func() {
		It("returns true if error is the same", func() {
			err := errors.New("error message")
			Expect(err).To(MatchError("error message"))
		})
	})
})
