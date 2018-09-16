package model

type Result struct {
	Price float64 `json:"price"`
}

type Insurances struct {
	Life bool `json:"life"`
}

//Data contains all personal Data
type Data struct {
	Age int `json:"age"`
}

//Input contains all request data in a structured way
type Input struct {
	Insurances Insurances `json:"insurances"`
	Data       Data       `json:"data"`
}
