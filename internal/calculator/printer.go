package calculator

import (
	"fmt"
)

var substring [5]string = [5]string{"",
	" (7 days)",
	" (Average)",
	" (90 days)",
	" (365 days)"}

func substringInitializer() {
	substring[0] += (" (" + fmt.Sprint(serviceLevels[0].hoursPerDay) + " hours)")
}

func errorPrinter(err error) {
	fmt.Println(err)
}

func printMaximumDowntime(sl string) {
	substringInitializer()
	message := "With a " + sl + "% Service Level, the maximum downtime allowed is:"

	for index := range serviceLevels {
		message += ("\n" + serviceLevels[index].label + substring[index] + ": " + serviceLevels[index].calculateDowntimeString())
	}

	fmt.Println(message)

}

func printMonitoringFrequency(input Input) error {

	message := "\nWith a " + input.MTTR + " MTTR, " + fmt.Sprint(input.Incidents) + " average incident per time period and " +
		fmt.Sprint(input.ProbeFailures) + " failed probe to alert,\n" +
		"Here is the MINIMUM probing frequency is necessary to keep your Service Level inside these time periods: "

	fmt.Println(message)

	for index := range serviceLevels {
		if serviceLevels[index].testingFrequencyNecessary != "" {
			fmt.Println(serviceLevels[index].label + substring[index] + ": " + serviceLevels[index].testingFrequencyNecessary)
		}
	}

	return nil

}

func printReverse(outage string) {
	substringInitializer()
	fmt.Println("Service Level for", outage, "seconds of downtime:")

	for index := range serviceLevels {
		fmt.Println(serviceLevels[index].label + substring[index] + ": " + serviceLevels[index].availabilityToString())
	}
}
