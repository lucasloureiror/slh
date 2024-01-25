package main

import (
	"fmt"

	"github.com/lucasloureiror/slh/internal/calculator"
)

func main() {
	var input string

	fmt.Scanln(&input)

	calculator.Start(input)

}
