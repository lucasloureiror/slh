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

func Calculate(input *Input) error {
	populateServiceLevelData(&serviceLevels, input.HoursPerDay)
	sl, err := checkSLInput(&input.ServiceLevel)

	downtimeCalculator := downtimeCalculator{}

	if err != nil {
		return err
	}
	for index := range serviceLevels {
		serviceLevels[index].calculator = &downtimeCalculator
		serviceLevels[index].data.availabilityInPercentage = sl
	}

	printMaximumDowntime(input.ServiceLevel)

	if input.MTTR != "" {
		err := calculateMonitoringFrequency(input.MTTR, input.Incidents, input.ProbeFailures)
		if err != nil {
			errorPrinter(err)
			return err
		}
		printMonitoringFrequency(*input)
	}
	return nil
}

func (d *downtimeCalculator) calculate(data *serviceLevelData) {
	data.downtimeInSeconds = data.totalTimePeriod * (1 - data.availabilityInPercentage/100)
}

func (d *downtimeCalculator) print(data *serviceLevelData) string {
	data.availabilityInPercentage = data.availabilityInPercentage / 100
	downtimePercentage := 1 - data.availabilityInPercentage
	data.downtimeInSeconds = downtimePercentage * data.totalTimePeriod
	downtimeString := convert.SecondsToTimeString(int(data.downtimeInSeconds))
	return downtimeString
}
