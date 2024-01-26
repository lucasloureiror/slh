package calculator

import (
	"github.com/lucasloureiror/slh/internal/models"
)

func weeklyCalculator(sla float64) {
	const secondsPerWeek = 604800
	sla = sla / 100
	downtime := 1 - sla

	downtime = downtime * secondsPerWeek

	downtimeString := convertSecondsToTimeString(int(downtime))

	outage := models.OutageAllowed{Seconds: int(downtime), TimeFrame: "Weekly: ", ResultInTimeString: downtimeString}
	outages = append(outages, outage)

}
