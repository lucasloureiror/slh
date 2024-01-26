package calculator

import (
	"github.com/lucasloureiror/slh/internal/models"
)

var outages []models.OutageAllowed

func Start(input *models.Input) {
	slaCalculator(&input.SLA)
	printDowtime(input.SLA)

	if input.MTTR != "" {
		calculateMonitoringFrequency(input.MTTR)
		printMonitoringFrequency(input.MTTR)
	}

}

func slaCalculator(input *string) {

	sla := checkInput(input)
	dailyCalculator(sla)
	weeklyCalculator(sla)
	monthlyCalculator(sla)
	quartelyCalculator(sla)
	yearlyCalculator(sla)

}
