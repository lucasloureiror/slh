/*
Copyright 2024 github.com/lucasloureiror/slh maintainers

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
		return 1, fmt.Errorf("invalid Service Level")
	}

	if sla < 16 && sla > 1 {
		*input = "99."
		if sla < 3 {
			*input = "99"
		}

		for i := 0; i < int(sla)-2; i++ {
			*input += "9"
		}
		sla = (1 - math.Pow(10, -sla)) * 100
	}

	return sla, nil
}
