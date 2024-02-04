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

var serviceLevels [5]serviceLevel

func populateServiceLevelData(serviceLevels *[5]serviceLevel, hoursPerDay int) {

	var labels = [5]string{"Daily", "Weekly", "Monthly", "Quartely", "Yearly"}
	var timePeriodInSeconds = [5]float64{
		1 * float64(hoursPerDay) * 60 * 60,
		7 * float64(hoursPerDay) * 60 * 60,
		30.44 * float64(hoursPerDay) * 60 * 60,
		90 * float64(hoursPerDay) * 60 * 60,
		365 * float64(hoursPerDay) * 60 * 60,
	}

	for index := range serviceLevels {
		serviceLevels[index].label = labels[index]
		serviceLevels[index].totalTimePeriod = timePeriodInSeconds[index]
		serviceLevels[index].hoursPerDay = hoursPerDay
	}

}
