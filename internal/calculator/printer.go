package calculator

import "fmt"

func printDowtime(SLA string) {
	fmt.Println(SLA+"%", "availability represents the following maximum allowable downtime to meet your Service Level:")
	for _, outage := range outages {
		fmt.Println(outage.TimeFrame + outage.ResultInTimeString)
	}

}

func printMonitoringFrequency(mtr string) {
	fmt.Println("\nWith a " + mtr + " MTTR, the following MINIMUM probing frequency is necessary (with 1 failed probe to alert) to meet your Service Level inside these time frames with ONE INCIDENT:")
	for _, outage := range outages {
		if outage.TestingFrequencyNecessary != "" {
			fmt.Println(outage.TimeFrame + outage.TestingFrequencyNecessary)
		}
	}

}
