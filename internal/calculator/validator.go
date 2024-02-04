package calculator

import (
	"fmt"
	"math"
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

	if sla < 16 && sla > 1 {
		*input = fmt.Sprint(sla) + " nines or 99."
		for i := 0; i < int(sla)-2; i++ {
			*input += "9"
		}
		sla = (1 - math.Pow(10, -sla)) * 100
	}

	return sla, nil
}
