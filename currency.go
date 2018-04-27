package numeric

import (
	"fmt"
	"strconv"
	"strings"
)

// USDToCents forwards to TryUSDToCents but panics in the case of a non-nil error.
func USDToCents(dollarAmount string) uint64 {
	cents, err := TryUSDToCents(dollarAmount)
	if err != nil {
		panic(err)
	}
	return cents
}

// TryUSDToCents accepts a string in the form "$77,888.99" and returns the amount in USD cents as a uint64.
// It returns an error if the given dollar amount is invalid in any way. See unit tests for examples.
func TryUSDToCents(dollarAmount string) (amountInCents uint64, err error) {
	dollarAmount = strings.TrimSpace(dollarAmount)

	if !isValidDollarAmount(dollarAmount) {
		return 0, fmt.Errorf("Invalid dollar amount: [%s] (Must be in the form: [$77,888.99])", dollarAmount)
	}
	dollarAmount = strings.Replace(dollarAmount, "$", "", 1)
	dollarAmount = strings.Replace(dollarAmount, ".", "", 1)
	dollarAmount = strings.Replace(dollarAmount, ",", "", -1)
	for strings.HasPrefix(dollarAmount, "0") && len(dollarAmount) > 1 {
		dollarAmount = dollarAmount[1:]
	}
	return strconv.ParseUint(dollarAmount, 10, 64)
}
func isValidDollarAmount(dollarAmount string) bool {
	const dollarSignIndex = 0
	var decimalPointIndex = len(dollarAmount) - 3
	for c, character := range dollarAmount {
		switch {
		case c == dollarSignIndex && character != '$':
			return false
		case c == decimalPointIndex && character != '.':
			return false
		case character == ',' && !isValidUSDSeparatorIndex(decimalPointIndex, c):
			return false
		case !isAllowedUSDCharacter(character):
			return false
		}
	}
	return true
}
func isAllowedUSDCharacter(r rune) bool {
	const allowed = "$,.0123456789"
	return strings.ContainsRune(allowed, r)
}
func isValidUSDSeparatorIndex(decimal, comma int) bool {
	return decimal > comma && (decimal - comma) % len(",000") == 0
}
