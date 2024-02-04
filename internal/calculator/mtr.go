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
