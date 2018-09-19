package insurance

import (
	"testing"

	"github.com/c0d1ngm0nk3y/sample-application-go/pkg/model"
	. "github.com/onsi/gomega"
)

func TestDefaultLife(t *testing.T) {
	g := NewGomegaWithT(t)

	data := model.Data{}

	result := CalculateLife(data)

	g.Expect(result).To(Equal(100.0))
}
