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

import (
	"strings"
	"testing"
)

func TestReverseCalculate(t *testing.T) {
	s := serviceLevel{data: serviceLevelData{totalTimePeriod: 86400, downtimeInSeconds: 8.0},
		calculator: &ReverseCalculator{},
	}
	i := Input{TotalOutageTime: "8s"}
	s.calculator.calculate(&s.data, &i)
	expectedSubstring := "99.99"
	got := s.calculator.print(&s.data)
	expected := strings.Contains(got, expectedSubstring)

	if expected != true {
		t.Errorf("Expected to contain %s, got %s", expectedSubstring, got)
	}
}
