package insurance

import "github.com/c0d1ngm0nk3y/sample-application-go/pkg/model"

//CalculateLife calculates the price of a life insurance for the given data
func CalculateLife(data model.Data) float64 {
	costs := 100.0

	if data.Age != 18 {
		costs = costs + 50
	}

	return costs
}
