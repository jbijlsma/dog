package dog

import (
	"testing"
)

var yearsTests = []struct {
	humanYears  int
	dogYears int
}{
	{1, 7},
	{7, 49},
	{10, 70},
}

func TestYears(t *testing.T) {
	for _, tt := range yearsTests {
		t.Run(string(tt.humanYears), func(t *testing.T) {
			got := Years(tt.humanYears)
			if got != tt.dogYears {
				t.Errorf("got %q, want %q", got, tt.dogYears)
			}
		})
	}
}