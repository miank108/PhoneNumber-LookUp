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

// Regex to validate only digits and optional + prefix
var phoneRegex = regexp.MustCompile(`^\+?\d+$`)

func ParsePhoneNumber(raw string, fallbackCC string) (PhoneResponse, error) {
	resp := PhoneResponse{}

	// Remove all spaces and normalize
	normalized := strings.ReplaceAll(raw, " ", "")

	// Validate format
	if !phoneRegex.MatchString(normalized) {
		resp.Error = map[string]string{"format": "invalid phone number format"}
		return resp, nil
	}

	// Max length validation
	if len(normalized) > 15 {
		resp.Error = map[string]string{"phoneNumber": "phone number exceeds maximum allowed length (15 digits)"}
		return resp, nil
	}

	// Ensure + prefix for consistency
	if !strings.HasPrefix(normalized, "+") {
		normalized = "+" + normalized
	}
	resp.PhoneNumber = normalized

	// Extract country code by trying different lengths
	withoutPlus := normalized[1:]
	// Try country codes from longest to shortest to avoid partial matches
	for i := 3; i >= 1; i-- {
		if len(withoutPlus) < i {
			continue
		}
		cc := withoutPlus[:i]
		if iso, ok := data.CountryDialMap[cc]; ok {
			resp.CountryCode = iso

			// Get the remaining number after country code
			remaining := withoutPlus[i:]

			// For numbers long enough, take first 3 digits as area code
			if len(remaining) > 3 {
				resp.AreaCode = remaining[:3]
				resp.LocalPhoneNumber = remaining[3:]
			} else {
				resp.LocalPhoneNumber = remaining
			}
			return resp, nil
		}
	}

	// If no country code found, check fallback
	if fallbackCC == "" {
		resp.Error = map[string]string{"countryCode": "required value is missing"}
		return resp, nil
	}

	// Validate fallback country code
	for _, iso := range data.CountryDialMap {
		if strings.EqualFold(fallbackCC, iso) {
			resp.CountryCode = iso
			if len(withoutPlus) > 3 {
				resp.AreaCode = withoutPlus[:3]
				resp.LocalPhoneNumber = withoutPlus[3:]
			} else {
				resp.LocalPhoneNumber = withoutPlus
			}
			return resp, nil
		}
	}

	resp.Error = map[string]string{"countryCode": "invalid country code"}
	return resp, nil

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
