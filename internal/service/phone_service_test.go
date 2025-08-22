package service

import (
	"testing"
)

func TestParsePhoneNumber_ValidWithCountryCode(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedCC  string
		expectedAC  string
		expectedLoc string
	}{
		{
			name:        "US number with +",
			input:       "+12125690123",
			expectedCC:  "US",
			expectedAC:  "212",
			expectedLoc: "5690123",
		},
		{
			name:        "Mexico with spaces",
			input:       "+52 631 3118150",
			expectedCC:  "MX",
			expectedAC:  "631",
			expectedLoc: "3118150",
		},
		{
			name:        "Spain without +",
			input:       "34915872200",
			expectedCC:  "ES",
			expectedAC:  "915",
			expectedLoc: "872200",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParsePhoneNumber(tc.input, "")
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.CountryCode != tc.expectedCC {
				t.Errorf("expected countryCode %s, got %s", tc.expectedCC, got.CountryCode)
			}
			if got.AreaCode != tc.expectedAC {
				t.Errorf("expected areaCode %s, got %s", tc.expectedAC, got.AreaCode)
			}
			if got.LocalPhoneNumber != tc.expectedLoc {
				t.Errorf("expected localPhoneNumber %s, got %s", tc.expectedLoc, got.LocalPhoneNumber)
			}
		})
	}
}

func TestParsePhoneNumber_UseFallbackCountryCode(t *testing.T) {
	got, err := ParsePhoneNumber("915872200", "ES")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.CountryCode != "ES" {
		t.Errorf("expected ES, got %s", got.CountryCode)
	}
	if got.AreaCode != "915" {
		t.Errorf("expected areaCode 915, got %s", got.AreaCode)
	}
	if got.LocalPhoneNumber != "872200" {
		t.Errorf("expected localPhoneNumber 872200, got %s", got.LocalPhoneNumber)
	}
}

func TestParsePhoneNumber_InvalidFormat(t *testing.T) {
	invalidNumbers := []string{
		"351 21 094   2000", // multiple spaces
		"+12A3456789",       // letters
		"++1234567890",      // invalid double +
	}

	for _, num := range invalidNumbers {
		_, err := ParsePhoneNumber(num, "")
		if err == nil {
			t.Errorf("expected error for invalid format: %s", num)
		}
	}
}
