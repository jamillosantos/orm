package testingutils

import (
	"testing"

	"github.com/novln/macchiato"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func InitMacchiato(t *testing.T, description string) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	macchiato.RunSpecs(t, description)
}
