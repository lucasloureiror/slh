package calculator

import ()

var outages []OutageAllowed

func Start(input *Input) error {
	slaCalculator(&input.SLA, input.HoursPerDay)
	printDowtime(*input)

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

func slaCalculator(input *string, hoursPerDay int) {

	sla := checkInput(input)
	dailyCalculator(sla, hoursPerDay)
	weeklyCalculator(sla, hoursPerDay)
	monthlyCalculator(sla, hoursPerDay)
	quartelyCalculator(sla, hoursPerDay)
	yearlyCalculator(sla, hoursPerDay)
}
