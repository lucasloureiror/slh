package calculator

import (
	"strings"
	"testing"
)

func TestReverseCalculate(t *testing.T) {
	s := serviceLevel{totalTimePeriod: 86400}
	s.reverseCalculate(8)
	expectedSubstring := "99.99"
	got := s.availabilityToString()
	expected := strings.Contains(got, expectedSubstring)

	if expected != true {
		t.Errorf("Expected to contain %s, got %s", expectedSubstring, got)
	}
}
