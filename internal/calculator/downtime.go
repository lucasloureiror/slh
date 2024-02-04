package calculator

import (
	"github.com/lucasloureiror/slh/internal/pkg"
)

func Calculate(input *Input) error {
	populateServiceLevelData(&serviceLevels, input.HoursPerDay)
	sl, err := checkSLInput(&input.ServiceLevel)

	if err != nil {
		return err
	}
	for index := range serviceLevels {
		serviceLevels[index].availabilityInPercentage = sl
	}

	printMaximumDowntime(input.ServiceLevel)

	if input.MTTR != "" {
		err := calculateMonitoringFrequency(input.MTTR, input.Incidents, input.ProbeFailures)
		if err != nil {
			errorPrinter(err)
			return err
		}
		printMonitoringFrequency(*input)
	}
	return nil
}

func (s *serviceLevel) calculateDowntimeString() string {
	s.availabilityInPercentage = s.availabilityInPercentage / 100
	downtimePercentage := 1 - s.availabilityInPercentage
	s.downtimeInSeconds = downtimePercentage * s.totalTimePeriod
	downtimeString := pkg.ConvertSecondsToTimeString(int(s.downtimeInSeconds))
	return downtimeString
}
