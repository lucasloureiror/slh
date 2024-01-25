package calculator

import (
	"fmt"
	"github.com/lucasloureiror/slh/internal/models"
)

func yearlyCalculator(sla float64) {
	//Using a year with 365 days
	const secondsPerYear = 365 * 24 * 60 * 60
	sla = sla / 100
	downtime := 1 - sla

	downtime = downtime * secondsPerYear
	days := int(downtime) / 86400
	hours := int(downtime) / 3600
	minutes := (int(downtime) % 3600) / 60
	seconds := int(downtime) % 60

	downtimeString := ""
	if (days + hours + minutes + seconds) == 0 {
		downtimeString = "0s"
	} else {
		if days > 0 {
			downtimeString += fmt.Sprintf("%dd ", days)
		}
		if hours > 0 {
			downtimeString += fmt.Sprintf("%dh ", hours)
		}
		if minutes > 0 {
			downtimeString += fmt.Sprintf("%dm ", minutes)
		}
		if seconds > 0 {
			downtimeString += fmt.Sprintf("%ds ", seconds)
		}

	}
	outage := models.OutageAllowed{Seconds: int(downtime), TimeFrame: "Yearly (365 days): ", ResultInTimeString: downtimeString}
	outages = append(outages, outage)
}
