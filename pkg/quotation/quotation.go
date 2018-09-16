package quotation

import (
	"encoding/json"

	"github.com/c0d1ngm0nk3y/sample-application-go/pkg/insurance"
	"github.com/c0d1ngm0nk3y/sample-application-go/pkg/model"
)

//Do will check the data and calculate a quotation
func Do(in []byte) (float64, error) {
	var input model.Input
	err := json.Unmarshal(in, &input)

	costs := 0.0

	if input.Insurances.Life {
		costs = costs + insurance.CalculateLife(input.Data)
	}

	return costs, err
}
