package calculator

import (
	"github.com/lucasloureiror/slh/internal/pkg"
)

func weeklyCalculator(sla float64, hoursPerDay int) {
	secondsPerWeek := hoursPerDay * 7 * 60 * 60
	sla = sla / 100
	downtime := 1 - sla

	downtime = downtime * float64(secondsPerWeek)

	downtimeString := pkg.ConvertSecondsToTimeString(int(downtime))

	outage := OutageAllowed{Seconds: int(downtime), TimeFrame: "Weekly: ", ResultInTimeString: downtimeString}
	outages = append(outages, outage)

}
