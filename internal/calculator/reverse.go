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

	for index := range serviceLevels {
		serviceLevels[index].reverseCalculate(float64(seconds))
	}

	printReverse(outage)

}

func (s *serviceLevel) reverseCalculate(downtime float64) {

	s.downtimeInSeconds = downtime

	if s.totalTimePeriod > float64(downtime) {
		s.availabilityInPercentage = 100 - (float64(downtime) / s.totalTimePeriod * 100)
	}
}

func (s *serviceLevel) availabilityToString() string {
	return fmt.Sprintf("%.5f", s.availabilityInPercentage) + "%"
}
