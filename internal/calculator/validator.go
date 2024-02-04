package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

func checkSLInput(input *string) (float64, error) {
	*input = strings.Replace(*input, ",", ".", 1)
	*input = strings.Replace(*input, "%", "", 1)

	sla, err := strconv.ParseFloat(*input, 64)

	if err != nil {
		return 0, fmt.Errorf("invalid Service Level")
	}

	return sla, nil
}
