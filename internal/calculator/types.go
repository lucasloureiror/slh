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
