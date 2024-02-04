package pkg_test

import (
	"github.com/lucasloureiror/slh/internal/pkg"
	"testing"
)

func TestTimeStringToSeconds(t *testing.T) {

	timeString := []string{"1d", "1h", "1m", "1s", "1d1h1m1s", "1d1h1m", "1d1h1s", "1d1m1s", "1h1m1s", "1d1h", "1d1m", "1d1s", "1h1m", "1h1s", "1m1s"}
	seconds := []int{86400, 3600, 60, 1, 90061, 90060, 90001, 86461, 3661, 90000, 86460, 86401, 3660, 3601, 61}

	for i, v := range timeString {
		got, err := pkg.ConvertTimeStringToSeconds(v)

		if err != nil {
			t.Errorf("Error: %s", err)
		}

		if got != seconds[i] {
			t.Errorf("Expected %d, got %d", seconds[i], got)
		}
	}

}

func TestErrorTimeStringToSeconds(t *testing.T) {
	timeString := []string{"1d1h1m1", "1d1h1m1s1", "1d1h1m1s1s", "1d1h1m1s1s1s", "1d1h1m1s1s1s1s", "xd", "xm", "xh", "xs", "xdxmxmxs"}

	for _, v := range timeString {
		_, err := pkg.ConvertTimeStringToSeconds(v)

		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	}
}
