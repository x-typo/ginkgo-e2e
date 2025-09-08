package tests

import (
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/x-typ/ginkgo-e2e/internal/clients"
	"github.com/x-typ/ginkgo-e2e/internal/models/login"
	"github.com/x-typ/ginkgo-e2e/internal/services"
)

var _ = Describe("User Login Endpoint -", func() {
	var apiClient *resty.Client

	BeforeEach(func() {
		apiClient = clients.NewApiClient()
	})

	It("should return a 200 OK status and a valid token", func() {

		email := os.Getenv("MAIN_USERNAME")
		password := os.Getenv("MAIN_PASSWORD")

		Expect(email).NotTo(BeEmpty(), "MAIN_USERNAME environment variable not set")
		Expect(password).NotTo(BeEmpty(), "MAIN_PASSWORD environment variable not set")

		resp, err := services.LoginUser(apiClient, email, password)

		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode()).To(Equal(http.StatusOK))

		loginResponse := resp.Result().(*login.LoginResponse)

		Expect(loginResponse.Success).To(BeTrue())
		Expect(loginResponse.Message).To(Equal("Login successful"))
		Expect(loginResponse.Data.Email).To(Equal("x-typo@outlook.com"))
		Expect(loginResponse.Data.Token).NotTo(BeEmpty())
	})
})
