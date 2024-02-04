package calculator

import (
	"testing"
)

func TestCalculateDowntimeString(t *testing.T) {
	s := serviceLevel{totalTimePeriod: 86400, availabilityInPercentage: 99.99}
	expected := "8s"
	if s.calculateDowntimeString() != expected {
		t.Errorf("Expected %s, got %s", expected, s.calculateDowntimeString())
	}
}
