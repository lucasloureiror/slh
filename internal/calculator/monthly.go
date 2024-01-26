package calculator

import (
	"github.com/lucasloureiror/slh/internal/models"
)

func monthlyCalculator(sla float64) {
	//Using an average month with 30.44 days
	const averageSecondsPerMonth = 30.44 * 24 * 60 * 60
	sla = sla / 100
	downtime := 1 - sla

	downtime = downtime * averageSecondsPerMonth

	downtimeString := convertSecondsToTimeString(int(downtime))

	outage := models.OutageAllowed{Seconds: int(downtime), TimeFrame: "Monthly (Average): ", ResultInTimeString: downtimeString}
	outages = append(outages, outage)

}
