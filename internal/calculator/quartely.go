package calculator

import (
	"github.com/lucasloureiror/slh/internal/models"
)

func quartelyCalculator(sla float64) {
	//Using a quarter with 90 days
	const secondsPerQuarter = 90 * 24 * 60 * 60
	sla = sla / 100
	downtime := 1 - sla

	downtime = downtime * secondsPerQuarter

	downtimeString := convertSecondsToTimeString(int(downtime))

	outage := models.OutageAllowed{Seconds: int(downtime), TimeFrame: "Quartely (90 days): ", ResultInTimeString: downtimeString}
	outages = append(outages, outage)

}
