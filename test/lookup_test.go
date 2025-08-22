package test

import (
	"phonenumberlookup/lookup"
	"testing"
)

func TestLookupPhoneNumber(t *testing.T) {
	// Test valid phone number
	validPhone := "123-456-7890"
	validResult, err := lookup.LookupPhoneNumber(validPhone)
	validExpected := "Found: Alice, Country: USA, Type: mobile"
	if err != nil {
		t.Errorf("Unexpected error for valid phone: %v", err)
	}
	if validResult.String() != validExpected {
		t.Errorf("Expected '%s', got '%s'", validExpected, validResult.String())
	}

	// Test another valid phone number
	anotherValid := "999-090-8888"
	anotherResult, err := lookup.LookupPhoneNumber(anotherValid)
	anotherExpected := "Found: Bob, Country: Canada, Type: landline"
	if err != nil {
		t.Errorf("Unexpected error for valid phone: %v", err)
	}
	if anotherResult.String() != anotherExpected {
		t.Errorf("Expected '%s', got '%s'", anotherExpected, anotherResult.String())
	}
	// Test not found phone number
	notFoundPhone := "555-555-5555"
	notFoundResult, err := lookup.LookupPhoneNumber(notFoundPhone)
	if err == nil || err.Error() != "Phone number not found" {
		t.Errorf("Expected error 'Phone number not found', got '%v'", err)
	}
	if notFoundResult != nil {
		t.Errorf("Expected nil result for not found phone, got '%v'", notFoundResult)
	}

	// Test invalid phone number
	invalidPhone := "1234567890"
	invalidResult, err := lookup.LookupPhoneNumber(invalidPhone)
	if err == nil || err.Error() != "Invalid phone number format" {
		t.Errorf("Expected error 'Invalid phone number format', got '%v'", err)
	}
	if invalidResult != nil {
		t.Errorf("Expected nil result for invalid phone, got '%v'", invalidResult)
	}
}
