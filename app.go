package main

import (
	"fmt"

	"github.com/lucasloureiror/sla/internal/calculator"
)

func main() {
	var input string

	fmt.Scanln(&input)

	calculator.Start(input)

}
