package handler

import (
	"encoding/json"
	"net/http"
	"phonenumberlookup/internal/service"
)

func PhoneHandler(w http.ResponseWriter, r *http.Request) {
	phoneNumber := r.URL.Query().Get("phoneNumber")
	countryCode := r.URL.Query().Get("countryCode")

	if phoneNumber == "" {
		http.Error(w, "phoneNumber is required", http.StatusBadRequest)
		return
	}

	result, err := service.ParsePhoneNumber(phoneNumber, countryCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
