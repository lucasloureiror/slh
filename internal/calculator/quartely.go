package calculator

import "github.com/lucasloureiror/slh/internal/pkg"

func quartelyCalculator(sla float64, hoursPerDay int) {
	//Using a quarter with 90 days
	secondsPerQuarter := 90 * hoursPerDay * 60 * 60
	sla = sla / 100
	downtime := 1 - sla

	downtime = downtime * float64(secondsPerQuarter)

	downtimeString := pkg.ConvertSecondsToTimeString(int(downtime))

	outage := OutageAllowed{Seconds: int(downtime), TimeFrame: "Quartely (90 days): ", ResultInTimeString: downtimeString}
	outages = append(outages, outage)

}
