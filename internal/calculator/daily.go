package calculator

import "fmt"

func dailyCalculator(sla float64) string {
	const secondsPerDay = 86400
	sla = sla / 100
	downtime := 1 - sla

	downtime = downtime * secondsPerDay
	hours := int(downtime) / 3600
	minutes := (int(downtime) % 3600) / 60
	seconds := int(downtime) % 60

	if (hours + minutes + seconds) == 0 {
		return "Daily: 0s"
	}
	downtimeString := "Daily: "
	if hours > 0 {
		downtimeString += fmt.Sprintf("%dh ", hours)
	}
	if minutes > 0 {
		downtimeString += fmt.Sprintf("%dm ", minutes)
	}
	if seconds > 0 {
		downtimeString += fmt.Sprintf("%ds", seconds)
	}

	return downtimeString

}
