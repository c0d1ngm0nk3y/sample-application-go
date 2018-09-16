package model

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

var example = []byte(`{"insurances":
						{"life": true}
					  , "data":
					    { "age": 21 }}`)

func TestUnmarshalComplete(t *testing.T) {
	g := NewGomegaWithT(t)

	var data Input
	err := json.Unmarshal(example, &data)

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(data.Insurances.Life).To(BeTrue())
	g.Expect(data.Data.Age).To(Equal(21))
}
