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
	"github.com/lucasloureiror/slh/internal/convert"
)

func Reverse(outage string, hoursPerDay int) {
	seconds, err := convert.TimeStringToSeconds(outage)
	if err != nil {
		errorPrinter(err)
		return
	}

	populateServiceLevelData(&serviceLevels, hoursPerDay)
	reverseCalculator := reverseCalculator{}

	for index := range serviceLevels {
		serviceLevels[index].data.downtimeInSeconds = float64(seconds)
		serviceLevels[index].calculator = &reverseCalculator
		serviceLevels[index].calculator.calculate(&serviceLevels[index].data)
	}

	printReverse(outage)

}

func (r *reverseCalculator) print(data *serviceLevelData) string {
	return fmt.Sprintf("%.5f", data.availabilityInPercentage) + "%"
}

func (r *reverseCalculator) calculate(data *serviceLevelData) {

	if data.totalTimePeriod > data.downtimeInSeconds {
		data.availabilityInPercentage = (100 - data.downtimeInSeconds/data.totalTimePeriod*100)
	}
}
