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
	"fmt"

	"github.com/lucasloureiror/slh/internal/convert"
)

func calculateMonitoringFrequency(mtr string, incidents int, probeFailures int) error {

	p := probeFrequencyCalculator{
		incidents:      incidents,
		probesFailures: probeFailures,
	}
	mttrInSeconds, err := convert.TimeStringToSeconds(mtr)

	if err != nil {
		return err
	}
	p.mttrInSeconds = mttrInSeconds * p.incidents

	impossibleToMonitor := true

	for i := range serviceLevels {
		sl := &serviceLevels[i]
		sl.calculator = &p
		sl.data.meanTimeToRecoveryInSeconds = p.mttrInSeconds
		if int(sl.data.downtimeInSeconds) > p.mttrInSeconds {
			sl.calculator.calculate(&sl.data)
			impossibleToMonitor = false
		}
	}

	if impossibleToMonitor {
		return errors.New("\nYou can't monitor this service with this MTTR without violating the Service Level")
	}

	return nil

}

func (p *probeFrequencyCalculator) calculate(data *serviceLevelData) {
	p.minimumFrequency = (data.downtimeInSeconds) - float64(data.meanTimeToRecoveryInSeconds)/(float64(p.incidents)*float64(p.probesFailures))
	fmt.Println("seconds: ", p.minimumFrequency)
	fmt.Println("probing frequency: ", p.print(data))
}

func (p *probeFrequencyCalculator) print(data *serviceLevelData) string {
	got := convert.SecondsToTimeString(int(p.minimumFrequency))
	fmt.Println("seconds: ", p.minimumFrequency, "got: ", got)
	return convert.SecondsToTimeString(int(p.minimumFrequency))
}
