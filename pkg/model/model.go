package model

import (
	"encoding/json"
	"fmt"
	"strconv"
)

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

func StringToInput(in []byte) Input {
	var input Input
	err := json.Unmarshal(in, &input)

	if err != nil {
		return Input{}
	} else {
		return input
	}
}

func StringToInputFlat(in []byte) Input {
	var raw map[string]string

	err := json.Unmarshal(in, &raw)
	if err != nil {
		fmt.Println(err.Error())
		return Input{}
	}

	age, _ := strconv.Atoi(raw["age"])
	data := Data{Age: age}

	life, _ := strconv.ParseBool(raw["life"])
	insurances := Insurances{Life: life}
	fmt.Printf("insuraces: %v\n", insurances)

	input := Input{Data: data, Insurances: insurances}
	return input
}
