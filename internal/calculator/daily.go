package calculator

import (
	"github.com/lucasloureiror/slh/internal/models"
)

func dailyCalculator(sla float64) {
	const secondsPerDay = 86400
	sla = sla / 100
	downtime := 1 - sla

	downtime = downtime * secondsPerDay

	downtimeString := convertSecondsToTimeString(int(downtime))

	outage := models.OutageAllowed{Seconds: int(downtime), TimeFrame: "Daily: ", ResultInTimeString: downtimeString}
	outages = append(outages, outage)

}
