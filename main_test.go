package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/c0d1ngm0nk3y/sample-application-go/pkg/model"
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

func TestEmptyStringQuotation(t *testing.T) {
	g := NewGomegaWithT(t)

	router := getRouter()

	body := []byte("")
	request, err := http.NewRequest("POST", "/quotation", bytes.NewReader(body))
	g.Expect(err).NotTo(HaveOccurred())

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	g.Expect(response.Code).To(Equal(http.StatusBadRequest))
}

func TestEmptyJsonQuotation(t *testing.T) {
	g := NewGomegaWithT(t)

	router := getRouter()

	body := []byte("{}")
	request, err := http.NewRequest("POST", "/quotation", bytes.NewReader(body))
	g.Expect(err).NotTo(HaveOccurred())

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	g.Expect(response.Code).To(Equal(http.StatusOK))

	var result model.Result
	err = json.Unmarshal([]byte(response.Body.String()), &result)
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(result.Price).To(Equal(0.0))
}

func TestQuotationWithoutBody(t *testing.T) {
	g := NewGomegaWithT(t)

	router := getRouter()

	request, err := http.NewRequest("POST", "/quotation", nil)
	g.Expect(err).NotTo(HaveOccurred())

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	g.Expect(response.Code).To(Equal(http.StatusInternalServerError))
}
