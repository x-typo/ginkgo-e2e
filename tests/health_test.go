package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my-api-test-framework/internal/endpoints"
	"my-api-test-framework/internal/models"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var _ = Describe("Health Check API", func() {
	It("should return 200 OK with success message", func() {
		// Example: using a real server URL or a test server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"success":true,"status":200,"message":"Successful Request"}`))
		}))
		defer server.Close()

		resp, err := http.Get(server.URL + endpoints.HealthCheckEndpoint)
		Expect(err).To(BeNil())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))

		var health models.HealthResponse
		err = json.NewDecoder(resp.Body).Decode(&health)
		Expect(err).To(BeNil())

		Expect(health.Success).To(BeTrue())
		Expect(health.Status).To(Equal(200))
		Expect(health.Message).To(Equal("Successful Request"))
	})
})
