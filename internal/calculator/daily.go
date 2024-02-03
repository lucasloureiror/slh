package calculator

import (
	"fmt"

	"github.com/lucasloureiror/slh/internal/pkg"
)

func dailyCalculator(sla float64, hoursPerDay int) {
	secondsPerDay := hoursPerDay * 60 * 60
	sla = sla / 100
	downtime := 1 - sla

	downtime = downtime * float64(secondsPerDay)

	downtimeString := pkg.ConvertSecondsToTimeString(int(downtime))

	timeFrame := "Daily: "

	if hoursPerDay != 24 {
		timeFrame = "Daily (" + fmt.Sprint(hoursPerDay) + " hours): "
	}

	outage := OutageAllowed{Seconds: int(downtime), TimeFrame: timeFrame, ResultInTimeString: downtimeString}
	outages = append(outages, outage)

}
