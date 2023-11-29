package utils

import (
	"fmt"
	"strings"
)

type ConfirmationRequest string

func Confirmation(req ConfirmationRequest) bool {
	var s string

	fmt.Printf("%s (y/N): ", req)
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Println("no valid option.")
		return Confirmation(req)
	}

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {
		return true
	}

	return false
}

type Question string

func Ask(req Question) string {
	var s string

	fmt.Printf("%s ", req)
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Println(err)
	}

	return strings.TrimSpace(s)
}
