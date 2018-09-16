package quotation

import (
	"testing"

	"github.com/c0d1ngm0nk3y/sample-application-go/pkg/model"
	. "github.com/onsi/gomega"
)

func TestDoWithEmpty(t *testing.T) {
	g := NewGomegaWithT(t)

	price := Do(model.StringToInput([]byte("")))
	g.Expect(price).To(Equal(0.0))
}

func TestDoEmptyJsonShouldReturn0(t *testing.T) {
	g := NewGomegaWithT(t)

	price := Do(model.StringToInput([]byte("{}")))

	g.Expect(price).To(Equal(0.0))
}

func TestDoDefaultLifeCosts150(t *testing.T) {
	g := NewGomegaWithT(t)

	price := Do(model.StringToInput([]byte(`{"insurances" : { "life": true }}`)))

	g.Expect(price).To(Equal(150.0))
}

func TestDoDefaultLifeAge18Costs100(t *testing.T) {
	g := NewGomegaWithT(t)

	price := Do(model.StringToInput([]byte(`{"insurances" : { "life": true }, "data": {"age": 18}}`)))

	g.Expect(price).To(Equal(100.0))
}
