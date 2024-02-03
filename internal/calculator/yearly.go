package calculator

import (
	"github.com/lucasloureiror/slh/internal/pkg"
)

func yearlyCalculator(sla float64, hoursPerDay int) {
	//Using a year with 365 days
	secondsPerYear := 365 * hoursPerDay * 60 * 60
	sla = sla / 100
	downtime := 1 - sla

	downtime = downtime * float64(secondsPerYear)

	downtimeString := pkg.ConvertSecondsToTimeString(int(downtime))

	outage := OutageAllowed{Seconds: int(downtime), TimeFrame: "Yearly (365 days): ", ResultInTimeString: downtimeString}
	outages = append(outages, outage)
}
