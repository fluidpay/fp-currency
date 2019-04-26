package currency

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// StringToUint : returns a uint from a string value
func StringToUint(num string, alpha string) (uint, error) {
	ISO, err := GetISOFromAlpha(alpha)
	if err != nil {
		return 0, err
	}
	reg := regexp.MustCompile("[0-9]+")
	str := reg.FindAllString(num, -1)
	strJoin := strings.Join(str, "")
	fl, err := strconv.ParseFloat(strJoin, 64)
	if err != nil {
		return 0, ErrorUnableToFormatCurrencyFromString
	}
	if strings.Contains(num, ISO.Decimal) == true {
		split := strings.Split(num, ISO.Decimal)
		if len(split[1]) != ISO.Fraction {
			return 0, ErrorUnableToFormatCurrencyFromString
		}
		return uint(fl), nil
	}
	return uint(fl * math.Pow10(ISO.Fraction)), nil
}

// DisplayFull : returns a string with full currency formatting... "num" being the amount, "alpha" being the ISO three digit alphabetic code.
func DisplayFull(num uint, alpha string) (string, error) {
	ISO, err := GetISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}
	return FormatCurrency(num, ISO), nil
}

// DisplayWithAlpha : returns a string with full currency formatting with the symbol replaced by the ISO three digit alphabetic code... "num" being the amount, "alpha" being the ISO three digit alphabetic code.
func DisplayWithAlpha(num uint, alpha string) (string, error) {
	ISO, err := GetISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}
	currency := FormatCurrency(num, ISO)
	return SwapSymbolWithAlpha(currency, ISO.Symbol, ISO.Alpha), nil
}

// DisplayNoSymbol : returns a string with full currency formatting minus the ISO symbol... "num" being the amount, "alpha" being the ISO three digit alphabetic code.
func DisplayNoSymbol(num uint, alpha string) (string, error) {
	ISO, err := GetISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}
	currency := FormatCurrency(num, ISO)
	return RemoveSymbol(currency, ISO.Symbol), nil
}

// DisplayWithDecimal : returns a string with all currency formatting removed except decimal places... "num" being the amount, "alpha" being the ISO three digit alphabetic code.
func DisplayWithDecimal(num uint, alpha string) (string, error) {
	ISO, err := GetISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}
	currency := FormatCurrency(num, ISO)
	currency = RemoveSymbol(currency, ISO.Symbol)
	return RemoveDelimiter(currency, ISO.Delimiter), nil
}

// UintToString : returns a string with all currency formatting removed... "num" being the amount, "alpha" being the ISO three digit alphabetic code.
func UintToString(num uint, alpha string) (string, error) {
	_, err := GetISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(num), nil
}

// ListCurrencies : returns a list of currencies
func ListCurrencies(list []string) ([]Currency, error) {
	currencies := []Currency{}
	for _, v := range list {
		ISO, err := GetISOFromAlpha(v)
		if err != nil {
			return nil, err
		}
		currencies = append(currencies, ISO)
	}
	return currencies, nil
}
