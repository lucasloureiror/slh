package calculator

import "errors"

func calculateMonitoringFrequency(mtr string, incidents int, probeFailures int) error {

	mtrInSeconds, _ := MTRToSeconds(mtr)
	mtrInSeconds = mtrInSeconds * incidents

	impossibleToMonitor := true

	for i := range outages {
		if outages[i].Seconds > mtrInSeconds {
			Minimumfrequency := (outages[i].Seconds - mtrInSeconds) / (incidents * probeFailures)
			outages[i].TestingFrequencyNecessary = convertSecondsToTimeString(Minimumfrequency)
			impossibleToMonitor = false
		}
	}

	if impossibleToMonitor {
		return errors.New("\nYou can't monitor this service with this MTTR without violating the Service Level")
	}

	return nil

}
