package service

import (
	"errors"
	"phonenumberlookup/internal/data"
	"regexp"
	"strings"
)

type PhoneResponse struct {
	PhoneNumber      string            `json:"phoneNumber"`
	CountryCode      string            `json:"countryCode"`
	AreaCode         string            `json:"areaCode"`
	LocalPhoneNumber string            `json:"localPhoneNumber"`
	Error            map[string]string `json:"error,omitempty"`
}

// Regex to validate E.164 with spaces
var phoneRegex = regexp.MustCompile(`^\+?\d+( \d+)*$`)

func ParsePhoneNumber(raw string, fallbackCC string) (PhoneResponse, error) {

	resp := PhoneResponse{}
	normalized := strings.ReplaceAll(raw, " ", "")

	// Validate
	if !phoneRegex.MatchString(raw) {
		return resp, errors.New("invalid phone number format")
	}
	// Max length validation (E.164 max is 15 digits, but you may want to allow a bit more for spaces)
	digitCount := len(regexp.MustCompile(`\d`).FindAllString(normalized, -1))
	if digitCount > 15 {
		resp.Error = map[string]string{"phoneNumber": "phone number exceeds maximum allowed length (15 digits)"}
		return resp, nil
	}

	// Ensure leading +
	if !strings.HasPrefix(normalized, "+") {
		normalized = "+" + normalized
	}
	resp.PhoneNumber = normalized

	// Extract country code
	withoutPlus := normalized[1:]
	for i := 1; i <= 3 && i <= len(withoutPlus); i++ {
		cc := withoutPlus[:i]
		if iso, ok := data.CountryDialMap[cc]; ok {
			resp.CountryCode = iso
			if len(withoutPlus) > i+3 {
				resp.AreaCode = withoutPlus[i : i+3]
				resp.LocalPhoneNumber = withoutPlus[i+3:]
			}
			return resp, nil
		}
	}

	// Fallback to provided countryCode
	if fallbackCC == "" {
		return resp, errors.New("missing country code")
	}
	resp.CountryCode = strings.ToUpper(fallbackCC)
	if len(withoutPlus) > 3 {
		resp.AreaCode = withoutPlus[:3]
		resp.LocalPhoneNumber = withoutPlus[3:]
	} else {
		resp.AreaCode = withoutPlus
	}
	return resp, nil
}
