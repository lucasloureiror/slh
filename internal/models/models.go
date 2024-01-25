package models

type Input struct {
	HoursPerDay int
	MTTR         string
	SLA         string
}

type OutageAllowed struct {
	Seconds                   int
	TimeFrame                 string
	ResultInTimeString        string
	TestingFrequencyNecessary string
}
