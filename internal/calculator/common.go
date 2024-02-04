package calculator

var serviceLevels [5]serviceLevel

func populateServiceLevelData(serviceLevels *[5]serviceLevel, hoursPerDay int) {

	var labels = [5]string{"Daily", "Weekly", "Monthly", "Quartely", "Yearly"}
	var timePeriodInSeconds = [5]float64{
		1 * float64(hoursPerDay) * 60 * 60,
		7 * float64(hoursPerDay) * 60 * 60,
		30.44 * float64(hoursPerDay) * 60 * 60,
		90 * float64(hoursPerDay) * 60 * 60,
		365 * float64(hoursPerDay) * 60 * 60,
	}

	for index := range serviceLevels {
		serviceLevels[index].label = labels[index]
		serviceLevels[index].totalTimePeriod = timePeriodInSeconds[index]
		serviceLevels[index].hoursPerDay = hoursPerDay
	}

}
