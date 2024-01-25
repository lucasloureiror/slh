package calculator


func calculateMonitoringFrequency(mtr string) {
	mtrInSeconds, _ := MTRToSeconds(mtr)

	for i := range outages {
		if outages[i].Seconds > mtrInSeconds{
			Minimumfrequency := outages[i].Seconds - mtrInSeconds
			outages[i].TestingFrequencyNecessary = convertSecondsToTimeString(Minimumfrequency)
		}
	}

}

//DOWNTIME OF 100 minutes
//MTTR OF 30
//Frequency of at least 70 minutes

//Downtime of 60 minutes
//MTTR OF 20
//Frequency of at least 40 minutes