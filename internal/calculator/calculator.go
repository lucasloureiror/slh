package calculator

import (
	"github.com/lucasloureiror/slh/internal/models"
)

var outages []models.OutageAllowed

func Start(input *models.Input) error {
	slaCalculator(&input.SLA)
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

func slaCalculator(input *string) {

	sla := checkInput(input)
	dailyCalculator(sla)
	weeklyCalculator(sla)
	monthlyCalculator(sla)
	quartelyCalculator(sla)
	yearlyCalculator(sla)
}
