package lookup

import (
	"errors"
	"phonenumberlookup/models"
	"regexp"
)

func isValidPhoneNumber(number string) bool {
	// Simple validation: 123-456-7890 format
	re := regexp.MustCompile(`^\d{3}-\d{3}-\d{4}$`)
	return re.MatchString(number)
}

var phoneBook = map[string]models.PhoneNumber{
	"123-456-7890": {Number: "123-456-7890", Name: "Alice", Country: "USA", Type: "mobile"},
	"999-090-8888": {Number: "999-090-8888", Name: "Bob", Country: "Canada", Type: "landline"},
	"000-232-9999": {Number: "000-232-9999", Name: "Charlie", Country: "UK", Type: "mobile"},
	"555-123-4567": {Number: "555-123-4567", Name: "David", Country: "India", Type: "landline"},
}

func LookupPhoneNumber(number string) (*models.PhoneNumber, error) {
	if !isValidPhoneNumber(number) {
		return nil, errors.New("Invalid phone number format")
	}
	if info, found := phoneBook[number]; found {
		return &info, nil
	}
	return nil, errors.New("Phone number not found")
}
