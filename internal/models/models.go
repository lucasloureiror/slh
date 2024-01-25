package models

type Input struct {
	HoursPerDay   int
	MTTR          string
	SLA           string
	ProbeFailures int
	Incidents     int
}

type OutageAllowed struct {
	Seconds                   int
	TimeFrame                 string
	ResultInTimeString        string
	TestingFrequencyNecessary string
}
