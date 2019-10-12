package auth

import (
	"../base"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Auth", func() {
	var token string
	var err error
	BeforeSuite(func() {
		base.SetupServer(base.SERVER_TYPE_LOCALHOST)

	})
	It("should return token with no error", func() {
		token, err = GetToken("saranshmiglani@gmail.com", "private")
		Expect(token).NotTo(Equal(""))
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("should pass the token validation", func() {
		_, _, valid := ValidateToken(token)
		Expect(valid).To(Equal(true))
	})
	It("should fail the token validation", func() {
		_, _, valid := ValidateToken("")
		Expect(valid).To(Equal(false))
	})
	It("should be able to login with no error", func() {
		_, err = Login("saranshmiglani@gmail.com", "123456")
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("should show login error", func() {
		_, err = Login("unknown@gmail.com", "123456")
		Expect(err).Should(HaveOccurred())
	})
	It("should return bcrypt hash of the password with no error", func() {
		_, err = GetPasswordHash("123456")
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("should return error as string is shoer", func() {
		_, err = GetPasswordHash("1")
		Expect(err).Should(HaveOccurred())
	})
	It("should return error while getting bcrypt hash of the password", func() {
		_, err = GetPasswordHash("")
		Expect(err).Should(HaveOccurred())
	})
	It("should return no error with valid user", func() {
		_, err = getUserEntity("saranshmiglani@gmail.com")
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("should return error with invalid user", func() {
		_, err = getUserEntity("unknown@gmail.com")
		Expect(err).Should(HaveOccurred())
	})
	It("should return false with no error", func() {
		exists, err := checkIfUserExistsEntity("")
		Expect(exists).To(Equal(false))
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("should return false ", func() {
		Expect(userVerificationFromToken("", "")).To(Equal(false))
	})
	It("should return error with mongo down", func() {
		base.MONGO_BASE_URL = "asasa"
		exists, err := checkIfUserExistsEntity("")
		Expect(exists).To(Equal(false))
		Expect(err).Should(HaveOccurred())
	})
})
