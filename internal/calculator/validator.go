package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

func checkInput(input *string) float64 {
	*input = strings.Replace(*input, ",", ".", 1)
	*input = strings.Replace(*input, "%", "", 1)

	sla, err := strconv.ParseFloat(*input, 64)

	if err != nil {
		fmt.Println("Invalid input")
	}

	return sla
}

func MTRToSeconds(mtr string) (int, error) {

	parts := strings.Split(mtr, "h")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid format")
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}

	minutesStr := strings.TrimSuffix(parts[1], "m")

	minutes, err := strconv.Atoi(minutesStr)
	if err != nil {
		return 0, err
	}

	seconds := hours*3600 + minutes*60

	return seconds, nil
}
