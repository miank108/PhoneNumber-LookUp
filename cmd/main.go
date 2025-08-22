package main

import (
	"fmt"
	"net/http"

	"phonenumberlookup/internal/handler"
)

func main() {
	http.HandleFunc("/v1/phone-numbers", handler.PhoneHandler)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
