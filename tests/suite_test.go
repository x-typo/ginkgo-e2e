package tests

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGinkgoE2E(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgo E2E Automation Test Suite")
}

var _ = BeforeSuite(func() {
	fmt.Println("Attempting to load .env file for local development...")
	err := godotenv.Load("..\\.env")
	if err != nil {
		fmt.Println("No .env file found")
	} else {
		fmt.Println(".env file loaded successfully.")
	}
})
