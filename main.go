package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"phonenumberlookup/internal/service"
)

func main() {
	http.HandleFunc("/v1/phone-numbers", func(w http.ResponseWriter, r *http.Request) {
		phoneNumber := r.URL.Query().Get("phoneNumber")
		countryCode := r.URL.Query().Get("countryCode")
		resp, _ := service.ParsePhoneNumber(phoneNumber, countryCode)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
