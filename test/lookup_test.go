package test

import (
	"phonenumberlookup/lookup"
	"testing"
)

func TestLookupPhoneNumber(t *testing.T) {
	// Test valid phone number
	validPhone := "123-456-7890"
	validResult := lookup.LookupPhoneNumber(validPhone)
	validExpected := "Found: Alice"
	if validResult != validExpected {
		t.Errorf("Expected '%s', got '%s'", validExpected, validResult)
	}

	// Test another valid phone number
	anotherValid := "999-090-8888"
	anotherResult := lookup.LookupPhoneNumber(anotherValid)
	anotherExpected := "Found: Bob"
	if anotherResult != anotherExpected {
		t.Errorf("Expected '%s', got '%s'", anotherExpected, anotherResult)
	}
	// Test not found phone number
	notFoundPhone := "555-555-5555"
	notFoundResult := lookup.LookupPhoneNumber(notFoundPhone)
	notFoundExpected := "Phone number not found"
	if notFoundResult != notFoundExpected {
		t.Errorf("Expected '%s', got '%s'", notFoundExpected, notFoundResult)
	}

	// Test invalid phone number
	invalidPhone := "1234567890"
	invalidResult := lookup.LookupPhoneNumber(invalidPhone)
	invalidExpected := "Invalid phone number format"
	if invalidResult != invalidExpected {
		t.Errorf("Expected '%s', got '%s'", invalidExpected, invalidResult)
	}
}
