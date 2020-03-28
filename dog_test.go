package dog

import (
	"testing"
)

func TestYears(t *testing.T) {
	testCases := []struct {
		humanYears       int
		expectedDogYears int
	} {
		{
			humanYears:       1,
			expectedDogYears: 7,
		},
		{
			humanYears:       7,
			expectedDogYears: 49,
		},
		{
			humanYears:       10,
			expectedDogYears: 70,
		},
	}


	for _, tt := range testCases {
		t.Run(string(tt.humanYears), func(t *testing.T) {
			t.Parallel()
			got := Years(tt.humanYears)
			if got != tt.expectedDogYears {
				t.Errorf("got %q, want %q", got, tt.expectedDogYears)
			}
		})
	}
}