package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

func checkInput(input string) float64 {
	input = strings.Replace(input, ",", ".", 1)
	input = strings.Replace(input, "%", "", 1)

	sla, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println("Invalid input")
	}

	return sla
}
