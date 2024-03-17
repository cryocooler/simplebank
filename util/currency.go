package util

const (
	USD = "USD"
	EUR = "EUR"
	CHF = "CHF"
	SEK = "SEK"
)

// true if currency input is supproted
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CHF, SEK:
		return true
	}
	return false

}
