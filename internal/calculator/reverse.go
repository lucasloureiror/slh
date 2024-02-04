package calculator

import (
	"fmt"
	"github.com/lucasloureiror/slh/internal/pkg"
)

func Reverse(outage string, hoursPerDay int) {
	seconds, err := pkg.ConvertTimeStringToSeconds(outage)
	if err != nil {
		errorPrinter(err)
		return
	}

	populateServiceLevelData(&serviceLevels, hoursPerDay)

	for index := range serviceLevels {
		serviceLevels[index].reverseCalculate(float64(seconds))
	}

	printReverse(outage)

}

func (s *serviceLevel) reverseCalculate(downtime float64) {

	s.downtimeInSeconds = downtime

	if s.totalTimePeriod > float64(downtime) {
		s.availabilityInPercentage = 100 - (float64(downtime) / s.totalTimePeriod * 100)
	}
}

func (s *serviceLevel) availabilityToString() string {
	return fmt.Sprintf("%.5f", s.availabilityInPercentage) + "%"
}
