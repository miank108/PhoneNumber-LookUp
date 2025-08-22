package lookup

import (
	"regexp"
)

func isValidPhoneNumber(number string) bool {
	// Simple validation: 123-456-7890 format
	re := regexp.MustCompile(`^\d{3}-\d{3}-\d{4}$`)
	return re.MatchString(number)
}

var phoneBook = map[string]string{
	"123-456-7890": "Alice",
	"999-090-8888": "Bob",
	"000-232-9999": "Charlie",
	"555-123-4567": "David",
}

func LookupPhoneNumber(number string) string {
	if !isValidPhoneNumber(number) {
		return "Invalid phone number format"
	}
	if name, found := phoneBook[number]; found {
		return "Found: " + name
	}
	return "Phone number not found"
}
