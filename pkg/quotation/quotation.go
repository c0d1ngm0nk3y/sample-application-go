package quotation

import (
	"github.com/c0d1ngm0nk3y/sample-application-go/pkg/insurance"
	"github.com/c0d1ngm0nk3y/sample-application-go/pkg/model"
)

//Do will check the data and calculate a quotation
func Do(input model.Input) float64 {
	costs := 0.0

	count := 0
	if input.Insurances.Life {
		costs = costs + insurance.CalculateLife(input.Data)
		count = count + 1
	}
	if input.Insurances.Household {
		costs = costs + insurance.CalculateHousehold(input.Data)
		count = count + 1
	}
	if input.Insurances.Accident {
		costs = costs + insurance.CalculateAccident(input.Data)
		count = count + 1
	}
	if input.Insurances.Liability {
		costs = costs + insurance.CalculateLiability(input.Data)
		count = count + 1
	}

	costs = costs - insurance.CalculateBonus(costs, count)

	return costs
}
