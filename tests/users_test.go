package tests

import (
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/x-typ/ginkgo-e2e/internal/clients"
	"github.com/x-typ/ginkgo-e2e/internal/models/auth"
	"github.com/x-typ/ginkgo-e2e/internal/models/shared"
	"github.com/x-typ/ginkgo-e2e/internal/services"
)

var _ = Describe("User -", func() {
	var (
		apiClient *resty.Client
		email     string
		password  string
	)

	BeforeEach(func() {
		apiClient = clients.NewApiClient()
		email = os.Getenv("MAIN_USERNAME")
		password = os.Getenv("MAIN_PASSWORD")

		Expect(email).NotTo(BeEmpty(), "MAIN_USERNAME environment variable not set")
		Expect(password).NotTo(BeEmpty(), "MAIN_PASSWORD environment variable not set")
	})

	Context("when logging in", func() {
		It("should return a 200 OK status and a valid token with valid credentials", func() {
			By("making a login request to the API", func() {
				resp, err := services.LoginUser(apiClient, email, password)

				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode()).To(Equal(http.StatusOK))

				loginResponse := resp.Result().(*auth.LoginResponse)

				By("verifying the login response payload", func() {
					Expect(loginResponse.Success).To(BeTrue())
					Expect(loginResponse.Message).To(Equal("Login successful"))
					Expect(loginResponse.Data.Token).NotTo(BeEmpty())
					Expect(loginResponse.Data.Email).To(Equal(email))
				})
			})
		})
	})

	Context("when authenticated", func() {
		It("should successfully retrieve the user's profile information", func() {
			var token string

			By("authenticating the user to get a token", func() {
				loginResp, loginErr := services.LoginUser(apiClient, email, password)
				Expect(loginErr).NotTo(HaveOccurred())
				Expect(loginResp.StatusCode()).To(Equal(http.StatusOK), "Login failed, cannot proceed to get profile")

				loginData := loginResp.Result().(*auth.LoginResponse)
				token = loginData.Data.Token
				Expect(token).NotTo(BeEmpty(), "Token was empty after a successful login")
			})

			By("making a request to the user profile endpoint", func() {
				GinkgoWriter.Println("Auth Token being used:", token)
				profileResp, profileErr := services.GetUserProfile(apiClient, token)

				Expect(profileErr).NotTo(HaveOccurred())
				Expect(profileResp.StatusCode()).To(Equal(http.StatusOK))

				profileResponse := profileResp.Result().(*auth.UserProfileInfoResponse)

				By("verifying the profile response payload", func() {
					Expect(profileResponse.Success).To(BeTrue())
					Expect(profileResponse.Message).To(Equal("Profile successful"))
					Expect(profileResponse.Data.ID).NotTo(BeEmpty())
					Expect(profileResponse.Data.Name).To(Equal("Tyson"))
					Expect(profileResponse.Data.Email).To(Equal(email))
					Expect(profileResponse.Data.Phone).To(Equal("5556667777"))
					Expect(profileResponse.Data.Company).To(Equal("Test Company"))
				})
			})
		})
	})

	Context("when logging out", func() {
		var token string

		BeforeEach(func() {
			By("authenticating the user to get a token", func() {
				loginResp, loginErr := services.LoginUser(apiClient, email, password)
				Expect(loginErr).NotTo(HaveOccurred())
				Expect(loginResp.StatusCode()).To(Equal(http.StatusOK))
				token = loginResp.Result().(*auth.LoginResponse).Data.Token
				Expect(token).NotTo(BeEmpty())
			})
		})

		It("should successfully log out the user and invalidate the token", func() {
			By("making a request to the logout endpoint", func() {
				logoutResp, logoutErr := services.LogoutUser(apiClient, token)

				Expect(logoutErr).NotTo(HaveOccurred())
				Expect(logoutResp.StatusCode()).To(Equal(http.StatusOK))

				logoutResponse := logoutResp.Result().(*shared.BaseResponse)

				Expect(logoutResponse.Success).To(BeTrue())
				Expect(logoutResponse.Message).To(Equal("User has been successfully logged out"))
			})

			By("verifying the token is no longer valid", func() {
				profileResp, profileErr := services.GetUserProfile(apiClient, token)

				Expect(profileErr).NotTo(HaveOccurred())
				Expect(profileResp.StatusCode()).To(Equal(http.StatusUnauthorized))
			})
		})
	})
})
