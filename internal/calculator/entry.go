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
	"github.com/lucasloureiror/slh/internal/convert"
)

func Start(calculator Calculator, input Input) error {
	populateServiceLevelData(&serviceLevels, input.HoursPerDay)

	for i := range serviceLevels {
		serviceLevels[i].calculator = calculator
		serviceLevels[i].calculator.calculate(&serviceLevels[i].data, input)
	}

	output(calculator, input)

	if input.MTTR != "" {
		err := startMonitoringFrequency(input)
		if err != nil {
			return err
		}
	}

	return nil

}

func startMonitoringFrequency(input Input) error {
	mttrInSeconds, err := convert.TimeStringToSeconds(input.MTTR)
	if err != nil {
		errorPrinter(err)
		return err
	}

	mttrInSeconds = mttrInSeconds * input.Incidents

	p := &ProbeFrequencyCalculator{}

	impossibleToMonitor := true

	for i := range serviceLevels {
		serviceLevels[i].calculator = p
		serviceLevels[i].data.meanTimeToRecoveryInSeconds = mttrInSeconds
		if int(serviceLevels[i].data.downtimeInSeconds) > serviceLevels[i].data.meanTimeToRecoveryInSeconds {
			serviceLevels[i].calculator.calculate(&serviceLevels[i].data, input)
			impossibleToMonitor = false
		}
	}

	if impossibleToMonitor {
		return errors.New("\nYou can't monitor this service with this MTTR without violating the Service Level")
	}

	printMonitoringFrequency(input)

	return nil

}
