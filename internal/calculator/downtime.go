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
	"github.com/lucasloureiror/slh/internal/convert"
)

func (d DowntimeCalculator) calculate(data *serviceLevelData, input Input) {
	sl, err := checkSLInput(&input.ServiceLevel)
	if err != nil {
		errorPrinter(err)
		return
	}
	data.availabilityInPercentage = sl
	data.downtimeInSeconds = data.totalTimePeriod * (1 - data.availabilityInPercentage/100)
}

func (d DowntimeCalculator) print(data *serviceLevelData) string {
	data.availabilityInPercentage = data.availabilityInPercentage / 100
	downtimePercentage := 1 - data.availabilityInPercentage
	data.downtimeInSeconds = downtimePercentage * data.totalTimePeriod
	downtimeString := convert.SecondsToTimeString(int(data.downtimeInSeconds))
	return downtimeString
}
