package calculator

import (
	"fmt"
	"github.com/lucasloureiror/slh/internal/models"
)

func dailyCalculator(sla float64) {
	const secondsPerDay = 86400
	sla = sla / 100
	downtime := 1 - sla

	downtime = downtime * secondsPerDay
	hours := int(downtime) / 3600
	minutes := (int(downtime) % 3600) / 60
	seconds := int(downtime) % 60

	downtimeString := ""
	if (hours + minutes + seconds) == 0 {
		downtimeString = "0s"
	} else {
		if hours > 0 {
			downtimeString += fmt.Sprintf("%dh ", hours)
		}
		if minutes > 0 {
			downtimeString += fmt.Sprintf("%dm ", minutes)
		}
		if seconds > 0 {
			downtimeString += fmt.Sprintf("%ds", seconds)
		}
	}

	outage := models.OutageAllowed{Seconds: int(downtime), TimeFrame: "Daily: ", ResultInTimeString: downtimeString}
	outages = append(outages, outage)

}
