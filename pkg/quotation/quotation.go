package quotation

import (
	"github.com/c0d1ngm0nk3y/sample-application-go/pkg/insurance"
	"github.com/c0d1ngm0nk3y/sample-application-go/pkg/model"
)

//Do will check the data and calculate a quotation
func Do(input model.Input) float64 {
	costs := 0.0

	if input.Insurances.Life {
		costs = costs + insurance.CalculateLife(input.Data)
	}
	if input.Insurances.Household {
		costs = costs + insurance.CalculateHousehold(input.Data)
	}
	if input.Insurances.Accident {
		costs = costs + insurance.CalculateAccident(input.Data)
	}
	if input.Insurances.Liability {
		costs = costs + insurance.CalculateLiability(input.Data)
	}

	return costs
}
