package tests

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/x-typ/ginkgo-e2e/internal/clients"
	"github.com/x-typ/ginkgo-e2e/internal/models/health"
	"github.com/x-typ/ginkgo-e2e/internal/services"
)

var _ = Describe("Health -", func() {
	var apiClient *resty.Client

	BeforeEach(func() {
		apiClient = clients.NewApiClient()
	})

	It("return a 200 OK status with a success message", func() {
		resp, err := services.CheckHealth(apiClient)

		Expect(err).NotTo(HaveOccurred())
		Expect(resp.StatusCode()).To(Equal(http.StatusOK))

		By("verifying the response body payload", func() {
			responseBody := resp.Result().(*health.HealthResponse)
			Expect(responseBody.Success).To(BeTrue())
			Expect(responseBody.Message).To(Equal("Notes API is Running"))
		})
	})
})
