package calculator

import (
	"github.com/lucasloureiror/slh/internal/models"
)

var downtime models.Downtime

func Start(input string) {
	downtime.SLA = input
	slaCalculator(input)

}

func slaCalculator(input string) {

	sla := checkInput(input)
	downtime.Daily = dailyCalculator(sla)
	downtime.Weekly = weeklyCalculator(sla)
	downtime.Monthly = monthlyCalculator(sla)
	downtime.Quartely = quartelyCalculator(sla)
	downtime.Yearly = yearlyCalculator(sla)

	print()

}
