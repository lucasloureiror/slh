package calculator

func Start(calculator Calculator, input Input) {
	populateServiceLevelData(&serviceLevels, input.HoursPerDay)

}
