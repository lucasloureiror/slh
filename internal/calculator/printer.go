package calculator

import (
	"fmt"

	"github.com/lucasloureiror/slh/internal/models"
)

func errorPrinter(err error) {
	fmt.Println(err)
}

func printDowtime(input models.Input) error {
	fmt.Println(input.SLA+"%", "availability represents the following maximum allowable downtime to meet your Service Level:")
	for _, outage := range outages {
		fmt.Println(outage.TimeFrame + outage.ResultInTimeString)
	}

	return nil

}

func printMonitoringFrequency(input models.Input) error {

	message := "\nWith a " + input.MTTR + " MTTR, " + fmt.Sprint(input.Incidents) + " average incident per time period and " +
		fmt.Sprint(input.ProbeFailures) + " failed probe to alert\n" +
		"Here is the MINIMUM probing frequency is necessary to keep your Service Level inside these time periods: "

	fmt.Println(message)

	for _, outage := range outages {
		if outage.TestingFrequencyNecessary != "" {
			fmt.Println(outage.TimeFrame + outage.TestingFrequencyNecessary)
		}
	}

	return nil

}
