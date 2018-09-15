package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/gomega"
)

func TestPing(t *testing.T) {
	g := NewGomegaWithT(t)

	router := getRouter()

	request, err := http.NewRequest("GET", "/ping", nil)
	g.Expect(err).NotTo(HaveOccurred())

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	g.Expect(response.Body.String()).To(Equal("pong"))
}
