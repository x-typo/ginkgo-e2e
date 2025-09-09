package tests

import (
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/x-typ/ginkgo-e2e/internal/clients"
	"github.com/x-typ/ginkgo-e2e/internal/models/auth"
	"github.com/x-typ/ginkgo-e2e/internal/services"
)

var _ = Describe("Login -", func() {
	var apiClient *resty.Client

	BeforeEach(func() {
		apiClient = clients.NewApiClient()
	})

	It("return a 200 OK status and a valid token", func() {
		var email, password string

		By("getting credentials from environment variables", func() {
			email = os.Getenv("MAIN_USERNAME")
			password = os.Getenv("MAIN_PASSWORD")

			Expect(email).NotTo(BeEmpty(), "MAIN_USERNAME environment variable not set")
			Expect(password).NotTo(BeEmpty(), "MAIN_PASSWORD environment variable not set")
		})

		By("making a login request to endpoint", func() {
			resp, err := services.LoginUser(apiClient, email, password)

			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))

			loginResponse := resp.Result().(*auth.LoginResponse)

			By("verifying the response body payload", func() {
				Expect(loginResponse.Success).To(BeTrue())
				Expect(loginResponse.Message).To(Equal("Login successful"))
				Expect(loginResponse.Data.Email).To(Equal(email))
				Expect(loginResponse.Data.Token).NotTo(BeEmpty())
			})
		})
	})
})
