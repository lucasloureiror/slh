package calculator

import "fmt"

func monthlyCalculator(sla float64) string {
	//Using an average month with 30.44 days
	const averageSecondsPerMonth = 30.44 * 24 * 60 * 60
	sla = sla / 100
	downtime := 1 - sla

	downtime = downtime * averageSecondsPerMonth
	days := int(downtime) / 86400
	hours := int(downtime) / 3600
	minutes := (int(downtime) % 3600) / 60
	seconds := int(downtime) % 60

	downtimeString := "Monthly (Average): "
	if days > 0 {
		downtimeString += fmt.Sprintf("%dd", days)
	}
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
