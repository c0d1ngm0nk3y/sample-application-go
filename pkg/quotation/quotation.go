package quotation

import (
	"encoding/json"
)

type input struct {
}

//Do will check the data and calculate a quotation
func Do(in []byte) (float64, error) {
	var input input
	err := json.Unmarshal(in, &input)

	return 0.0, err
}
