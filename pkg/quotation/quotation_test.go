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

func TestDoDefaultLifeCosts150(t *testing.T) {
	g := NewGomegaWithT(t)

	result, err := Do([]byte(`{"insurances" : { "life": true }}`))
	g.Expect(err).NotTo(HaveOccurred())

	g.Expect(result).To(Equal(150.0))
}

func TestDoDefaultLifeAge18Costs100(t *testing.T) {
	g := NewGomegaWithT(t)

	result, err := Do([]byte(`{"insurances" : { "life": true }, "data": {"age": 18}}`))
	g.Expect(err).NotTo(HaveOccurred())

	g.Expect(result).To(Equal(100.0))
}
