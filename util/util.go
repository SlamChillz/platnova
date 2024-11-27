package util

import (
	"fmt"
	"strings"
)

func AddCurrencySymbol(symbol string, value string) string {
	if value == "0.00" {
		return ""
	}
	return fmt.Sprintf("%s%s", symbol, value)
}

// FormatNumber formats a float to a string with comma as a thousand separator
// and always ensures two decimal places.
func FormatNumber(num float64) string {
	// Format the float with two decimal places
	formatted := fmt.Sprintf("%.2f", num)

	// Split into whole number and decimal parts
	parts := strings.Split(formatted, ".")
	wholeNumber := parts[0]
	decimal := parts[1]

	// Add commas to the whole number part
	var withCommas strings.Builder
	length := len(wholeNumber)
	for i, char := range wholeNumber {
		withCommas.WriteRune(char)
		// Add a comma every three digits, except at the end
		if (length-i-1)%3 == 0 && i != length-1 {
			withCommas.WriteRune(',')
		}
	}

	// Combine the whole number with commas and the decimal part
	return withCommas.String() + "." + decimal
}
