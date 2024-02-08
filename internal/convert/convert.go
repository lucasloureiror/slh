/*
Copyright 2024 github.com/lucasloureiror/slh maintainers

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package convert

import (
	"fmt"
	"regexp"
	"strconv"
)

func SecondsToTimeString(timeInSeconds int) string {
	days := int(timeInSeconds) / 86400
	hours := (int(timeInSeconds) % 86400) / 3600
	minutes := ((int(timeInSeconds) % 86400) % 3600) / 60
	seconds := (((int(timeInSeconds) % 86400) % 3600) % 60)

	convertedString := ""
	if (days + hours + minutes + seconds) == 0 {
		convertedString = "0s"
	} else {
		if days > 0 {
			convertedString += fmt.Sprintf("%dd ", days)
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

func TimeStringToSeconds(mtr string) (int, error) {
	re := regexp.MustCompile(`^(?:(\d+)d)?(?:(\d+)h)?(?:(\d+)m)?(?:(\d+)s)?$`)
	matches := re.FindStringSubmatch(mtr)

	if matches == nil || len(matches) < 5 {
		return 0, fmt.Errorf("invalid time string format: %s", mtr)
	}

	var days, hours, minutes, seconds int
	var err error

	if matches[1] != "" {
		days, err = strconv.Atoi(matches[1])
		if err != nil {
			return 0, fmt.Errorf("invalid number of days: %s", matches[1])
		}
	}

	if matches[2] != "" {
		hours, err = strconv.Atoi(matches[2])
		if err != nil {
			return 0, fmt.Errorf("invalid number of hours: %s", matches[2])
		}
	}

	if matches[3] != "" {
		minutes, err = strconv.Atoi(matches[3])
		if err != nil {
			return 0, fmt.Errorf("invalid number of minutes: %s", matches[3])
		}
	}

	if matches[4] != "" {
		seconds, err = strconv.Atoi(matches[4])
		if err != nil {
			return 0, fmt.Errorf("invalid number of seconds: %s", matches[4])
		}
	}

	result := days*86400 + hours*3600 + minutes*60 + seconds
	return result, nil
}
