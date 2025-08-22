package main

import (
	"bufio"
	"fmt"
	"os"
	"phonenumberlookup/lookup"
	"strings"
)

func main() {
	fmt.Println("Phone Number Lookup")
	fmt.Print("Enter a phone number (format: 123-456-7890): ")
	reader := bufio.NewReader(os.Stdin)
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)
	result, err := lookup.LookupPhoneNumber(phone)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}
