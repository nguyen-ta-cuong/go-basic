package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	name := strings.ToUpper(customer)
	return fmt.Sprintf("Welcome to the Tech Palace, %s", name)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var borderLine strings.Builder
	for range numStarsPerLine {
		borderLine.WriteString("*")
	}
	return borderLine.String() + "\n" + welcomeMsg + "\n" + borderLine.String()
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	result := strings.Replace(oldMsg, "*", " ", -1)
	return strings.TrimSpace(result)
}
