package calculator

import "github.com/lucasloureiror/slh/internal/pkg"

func monthlyCalculator(sla float64, hoursPerDay int) {
	//Using an average month with 30.44 days
	averageSecondsPerMonth := 30.44 * float64(hoursPerDay) * 60 * 60
	sla = sla / 100
	downtime := 1 - sla

	downtime = downtime * averageSecondsPerMonth

	downtimeString := pkg.ConvertSecondsToTimeString(int(downtime))

	outage := OutageAllowed{Seconds: int(downtime), TimeFrame: "Monthly (Average): ", ResultInTimeString: downtimeString}
	outages = append(outages, outage)

}
