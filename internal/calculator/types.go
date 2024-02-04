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

type Input struct {
	HoursPerDay     int
	MTTR            string
	ServiceLevel    string
	TotalOutageTime string
	ProbeFailures   int
	Incidents       int
}

type serviceLevel struct {
	label                     string
	hoursPerDay               int
	availabilityInPercentage  float64
	downtimeInSeconds         float64
	totalTimePeriod           float64
	testingFrequencyNecessary string
}
