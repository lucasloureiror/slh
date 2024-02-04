package calculator

import (
	"errors"
	"github.com/lucasloureiror/slh/internal/pkg"
)

func calculateMonitoringFrequency(mtr string, incidents int, probeFailures int) error {

	mtrInSeconds, err := pkg.ConvertTimeStringToSeconds(mtr)

	if err != nil {
		return err
	}
	mtrInSeconds = mtrInSeconds * incidents

	impossibleToMonitor := true

	for i := range serviceLevels {
		if int(serviceLevels[i].downtimeInSeconds) > mtrInSeconds {
			Minimumfrequency := (int(serviceLevels[i].downtimeInSeconds) - mtrInSeconds) / (incidents * probeFailures)
			serviceLevels[i].testingFrequencyNecessary = pkg.ConvertSecondsToTimeString(Minimumfrequency)
			impossibleToMonitor = false
		}
	}

	if impossibleToMonitor {
		return errors.New("\nYou can't monitor this service with this MTTR without violating the Service Level")
	}

	return nil

}
