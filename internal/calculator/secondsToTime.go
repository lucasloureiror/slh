package calculator

import "fmt"

func convertSecondsToTimeString(timeInSeconds int) string{
	days := int(timeInSeconds) / 86400
	hours := int(timeInSeconds) / 3600
	minutes := (int(timeInSeconds) % 3600) / 60
	seconds := int(timeInSeconds) % 60

	convertedString := ""
	if (days + hours + minutes + seconds) == 0 {
		convertedString = "0s"
	} else {
		if days > 0 {
			convertedString += fmt.Sprintf("%dd", days)
		}
		if hours > 0 {
			convertedString += fmt.Sprintf("%dh ", hours)
		}
		if minutes > 0 {
			convertedString += fmt.Sprintf("%dm ", minutes)
		}
		if seconds > 0 {
			convertedString += fmt.Sprintf("%ds", seconds)
		}
		
	}

	return convertedString

}

