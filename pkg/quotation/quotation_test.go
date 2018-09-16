package quotation

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestDoWithEmpty(t *testing.T) {
	g := NewGomegaWithT(t)

	_, err := Do([]byte(""))
	g.Expect(err).To(HaveOccurred())
}

func TestDoEmptyJsonShouldReturn0(t *testing.T) {
	g := NewGomegaWithT(t)

	result, err := Do([]byte("{}"))
	g.Expect(err).NotTo(HaveOccurred())

	g.Expect(result).To(Equal(0.0))
}
