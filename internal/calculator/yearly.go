package calculator

import (
	"github.com/lucasloureiror/slh/internal/models"
)

func yearlyCalculator(sla float64) {
	//Using a year with 365 days
	const secondsPerYear = 365 * 24 * 60 * 60
	sla = sla / 100
	downtime := 1 - sla

	downtime = downtime * secondsPerYear

	downtimeString := convertSecondsToTimeString(int(downtime))

	outage := models.OutageAllowed{Seconds: int(downtime), TimeFrame: "Yearly (365 days): ", ResultInTimeString: downtimeString}
	outages = append(outages, outage)
}
