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
)

var substring = [5]string{"",
	" (7 days)",
	" (Average)",
	" (90 days)",
	" (365 days)"}

func substringInitializer() {
	substring[0] += (" (" + fmt.Sprint(serviceLevels[0].data.hoursPerDay) + " hours)")
}

func errorPrinter(err error) {
	fmt.Println(err)
}

func printMaximumDowntime(sl string) {
	substringInitializer()
	message := "With a " + sl + "% Service Level, the maximum downtime allowed is:"

	for index := range serviceLevels {
		message += ("\n" + serviceLevels[index].label + substring[index] + ": " + serviceLevels[index].calculator.print(&serviceLevels[index].data))
	}

	fmt.Println(message)

}

func printMonitoringFrequency(input Input) error {

	message := "\nWith a " + input.MTTR + " MTTR, " + fmt.Sprint(input.Incidents) + " average incident per time period and " +
		fmt.Sprint(input.ProbeFailures) + " failed probe to alert,\n" +
		"Here is the MINIMUM probing frequency is necessary to keep your Service Level inside these time periods: "

	fmt.Println(message)

	for index := range serviceLevels {
		testingFrequencyNecessary := serviceLevels[index].calculator.print(&serviceLevels[index].data)
		if testingFrequencyNecessary != "0s" {
			fmt.Println(serviceLevels[index].label + substring[index] + ": " + testingFrequencyNecessary)
		}
	}

	return nil

}

func printReverse(outage string) {
	substringInitializer()
	fmt.Println("Service Level for", outage, "of downtime:")

	for index := range serviceLevels {
		fmt.Println(serviceLevels[index].label + substring[index] + ": " + serviceLevels[index].calculator.print(&serviceLevels[index].data))
	}
}
