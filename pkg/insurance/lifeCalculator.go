package insurance

import "github.com/c0d1ngm0nk3y/sample-application-go/pkg/model"

//CalculateLife calculates the price of a life insurance for the given data
func CalculateLife(data model.Data) float64 {
	costs := 100.0

	if data.Age >= 30 {
		costs = costs + 50
	}

	if data.Male {
		costs = costs + 25
	}

	if data.Married || data.Kids >= 2 { //fix
		costs = costs - 25
	}

	return costs
}

func CalculateHousehold(data model.Data) float64 {
	costs := 75.0

	if data.Married {
		costs = costs + 50
	}

	return costs
}

func CalculateAccident(data model.Data) float64 {
	costs := 100.0

	if !data.Male {
		costs = costs + 25
	}

	if data.Married {
		costs = costs - 25
	}

	return costs
}

func CalculateLiability(data model.Data) float64 {
	costs := 250.0

	if data.Married {
		costs = costs + 50
	}
	costs = costs + (float64(data.Kids) * 25.0)

	return costs
}
