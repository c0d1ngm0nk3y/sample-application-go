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
	Life      bool `json:"life"`
	Household bool `json:"household"`
	Liability bool `json:"liability"`
	Accident  bool `json:"accident"`
}

//Data contains all personal Data
type Data struct {
	Age     int  `json:"age"`
	Kids    int  `json:"kids"`
	Married bool `json:"married"`
	Male    bool `json:"married"`
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
	kids, _ := strconv.Atoi(raw["kids"])
	married, _ := strconv.ParseBool(raw["married"])
	male, _ := strconv.ParseBool(raw["male"])
	data := Data{Age: age, Kids: kids, Married: married, Male: male}
	fmt.Printf("data: %v\n", data)

	life, _ := strconv.ParseBool(raw["life"])
	household, _ := strconv.ParseBool(raw["household"])
	liability, _ := strconv.ParseBool(raw["liability"])
	accident, _ := strconv.ParseBool(raw["accident"])
	insurances := Insurances{Life: life, Household: household, Liability: liability, Accident: accident}
	fmt.Printf("insuraces: %v\n", insurances)

	input := Input{Data: data, Insurances: insurances}
	return input
}
