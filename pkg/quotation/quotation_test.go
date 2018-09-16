package quotation

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestDoWithEmpty(t *testing.T) {
	g := NewGomegaWithT(t)

	err := Do([]byte(""))
	g.Expect(err).To(HaveOccurred())
}
