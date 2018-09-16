package quotation

import (
	"encoding/json"
)

type insurances struct {
	Life bool `json:"life"`
}

type data struct {
	Age int `json:"age"`
}

type input struct {
	Insurances insurances `json:"insurances"`
	Data       data       `json:"data"`
}

//Do will check the data and calculate a quotation
func Do(in []byte) (float64, error) {
	var input input
	err := json.Unmarshal(in, &input)

	costs := 0.0

	if input.Insurances.Life {
		costs = costs + 100

		if input.Data.Age != 18 {
			costs = costs + 50
		}
	}

	return costs, err
}
