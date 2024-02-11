package calculator

import (
	"testing"
)

func TestCalculateMinimumFrequency(t *testing.T) {
	p := ProbeFrequencyCalculator{}

	sl := serviceLevel{
		data: serviceLevelData{
			downtimeInSeconds:           6048.0,
			meanTimeToRecoveryInSeconds: 3600,
			totalTimePeriod:             604800},
		calculator: &p,
	}
	i := Input{
		ProbeFailures: 2,
		Incidents:     3}

	sl.calculator.calculate(&sl.data, i)

	got := sl.data.minimumFrequency
	expected := 408.0

	if got != expected {
		t.Errorf("Expected %f, got %f", expected, got)
	}
}

func TestPrintMinimumFrequency(t *testing.T) {
	p := ProbeFrequencyCalculator{}

	sl := serviceLevel{
		data: serviceLevelData{
			downtimeInSeconds:           6048.0,
			meanTimeToRecoveryInSeconds: 3600,
			totalTimePeriod:             604800},
		calculator: &p,
	}
	sl.calculator.calculate(&sl.data, Input{ProbeFailures: 2, Incidents: 3})
	got := sl.calculator.print(&sl.data)
	expected := "6m 48s"

	if got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}
