package dough

import (
	"reflect"
	"testing"
)

var TestStringToIntData = []struct {
	Num        string
	Alpha      string
	AllowLoose bool
	Output     interface{}
}{
	// Invalid values
	{"", "USD", false, ErrorInvalidStringFormat.Error()},
	{"     ", "USD", false, ErrorInvalidStringFormat.Error()},
	{"abcd", "USD", false, ErrorInvalidStringFormat.Error()},
	{"$5", "USA", false, ErrorInvalidISO.Error()},
	{"$0.0.5", "USD", false, ErrorInvalidStringFormat.Error()},
	{"$5.0", "USD", false, ErrorInvalidISOFractionMatch.Error()},
	{"$5.000", "USD", false, ErrorInvalidISOFractionMatch.Error()},

	//  Loose Validation tests
	{"$5.1", "USD", true, 510},
	{"$5.", "USD", true, 500},
	{"$5.001", "USD", true, 500},
	{"$5.101", "USD", true, 510},

	// Various Standard values
	{"$5", "USD", false, 500},
	{"$500", "USD", false, 50000},
	{"$-500", "USD", false, -50000},
	{"$05", "USD", false, 500},
	{"$0.05", "USD", false, 5},
	{"$5.52", "USD", false, 552},
	{"$0.00", "USD", false, 0},
	{"$0.01", "USD", false, 1},
	{"$0.10", "USD", false, 10},
	{"$1.00", "USD", false, 100},
	{"$10.00", "USD", false, 1000},
	{"$100.00", "USD", false, 10000},
	{"$1,000.00", "USD", false, 100000},
	{"$10,000.00", "USD", false, 1000000},
	{"$100,000.00", "USD", false, 10000000},
	{"$1,000,000.00", "USD", false, 100000000},

	// Problematic Numbers
	{"$538.92", "USD", false, 53892},
	{"$65.85", "USD", false, 6585},
	{"$17.99", "USD", false, 1799},
	{"538.92", "USD", false, 53892},
	{"65.85", "USD", false, 6585},
	{"17.99", "USD", false, 1799},
	{"$-538.92", "USD", false, -53892},
	{"$-65.85", "USD", false, -6585},
	{"$-17.99", "USD", false, -1799},
	{"-538.92", "USD", false, -53892},
	{"-65.85", "USD", false, -6585},
	{"-17.99", "USD", false, -1799},
	{"-$538.92", "USD", false, -53892},
	{"-$65.85", "USD", false, -6585},
	{"-$17.99", "USD", false, -1799},

	// Non USD
	{"$100.00,00", "ARS", false, 1000000},
	{"$10,000,000", "JPY", false, 10000000},
}

func TestStringToInt(t *testing.T) {
	// Test various structs
	for _, v := range TestStringToIntData {
		result, err := StringToInt(v.Num, v.Alpha, v.AllowLoose)
		if err != nil {
			if err.Error() != v.Output {
				t.Error("Expected:", v.Output, "Error:", err.Error())
			}
		} else if result != v.Output {
			t.Error("Expected:", v.Output, "Got:", result)
		}
	}

	// Test seeding numbers
	for _, v := range TestLargeNums {
		result, err := StringToInt(v.String2, "USD")
		if err != nil {
			t.Error(err)
		}
		if result != v.Integer {
			t.Error("Expected:", v.Integer, "Got:", result)
		}
	}
}

func BenchmarkStringToInt(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		StringToInt("12,345,678.99", "USD")
	}
}

func BenchmarkStringToIntLoose(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		StringToInt("12,345,678.99", "USD", true)
	}
}

var TestDisplayFullData = []struct {
	Amount int
	Alpha  string
	Output string
}{
	{0, "USA", ErrorInvalidISO.Error()},
	{0, "USD", "$0.00"},
	{1, "USD", "$0.01"},
	{10, "USD", "$0.10"},
	{100, "USD", "$1.00"},
	{1000, "USD", "$10.00"},
	{10000, "USD", "$100.00"},
	{100000, "USD", "$1,000.00"},
	{1000000, "USD", "$10,000.00"},
	{10000000, "USD", "$100,000.00"},
	{100000000, "USD", "$1,000,000.00"},
	{0, "AED", "0.00\u0625\u002E\u062F"},
	{1, "AED", "0.01\u0625\u002E\u062F"},
	{10, "AED", "0.10\u0625\u002E\u062F"},
	{100, "AED", "1.00\u0625\u002E\u062F"},
	{1000, "AED", "10.00\u0625\u002E\u062F"},
	{10000, "AED", "100.00\u0625\u002E\u062F"},
	{100000, "AED", "1,000.00\u0625\u002E\u062F"},
	{1000000, "AED", "10,000.00\u0625\u002E\u062F"},
	{10000000, "AED", "100,000.00\u0625\u002E\u062F"},
	{100000000, "AED", "1,000,000.00\u0625\u002E\u062F"},
	{0, "AED", "0.00\u0625\u002E\u062F"},
	{-1, "AED", "-0.01\u0625\u002E\u062F"},
	{-10, "AED", "-0.10\u0625\u002E\u062F"},
	{-100, "AED", "-1.00\u0625\u002E\u062F"},
	{-1000, "AED", "-10.00\u0625\u002E\u062F"},
	{-10000, "AED", "-100.00\u0625\u002E\u062F"},
	{-100000, "AED", "-1,000.00\u0625\u002E\u062F"},
	{-1000000, "AED", "-10,000.00\u0625\u002E\u062F"},
	{-10000000, "AED", "-100,000.00\u0625\u002E\u062F"},
	{-100000000, "AED", "-1,000,000.00\u0625\u002E\u062F"},
	{-0, "USD", "$0.00"},
	{-1, "USD", "$-0.01"},
	{-10, "USD", "$-0.10"},
	{-100, "USD", "$-1.00"},
	{-1000, "USD", "$-10.00"},
	{-10000, "USD", "$-100.00"},
	{-100000, "USD", "$-1,000.00"},
	{-1000000, "USD", "$-10,000.00"},
	{-10000000, "USD", "$-100,000.00"},
	{-100000000, "USD", "$-1,000,000.00"},
}

func TestDisplayFull(t *testing.T) {
	for _, v := range TestDisplayFullData {
		result, err := DisplayFull(v.Amount, v.Alpha)
		if err != nil {
			if err.Error() != v.Output {
				t.Error(err.Error())
			}
		} else if result != v.Output {
			t.Error(result)
		}
	}
}

var TestDisplayWithAlphaData = []struct {
	Amount int
	Alpha  string
	Output string
}{
	{0, "USA", ErrorInvalidISO.Error()},
	{0, "USD", "USD 0.00"},
	{-1, "USD", "USD -0.01"},
}

func TestDisplayWithAlpha(t *testing.T) {
	for _, v := range TestDisplayWithAlphaData {
		result, err := DisplayWithAlpha(v.Amount, v.Alpha)
		if err != nil {
			if err.Error() != v.Output {
				t.Error(err.Error())
			}
		} else if result != v.Output {
			t.Error(result)
		}
	}
}

var TestDisplayNoSymbolData = []struct {
	Num    int
	Alpha  string
	Output string
}{
	{0, "USA", ErrorInvalidISO.Error()},
	{0, "USD", "0.00"},
	{1, "USD", "0.01"},
	{10, "USD", "0.10"},
	{100, "USD", "1.00"},
	{1000, "USD", "10.00"},
	{10000, "USD", "100.00"},
	{100000, "USD", "1,000.00"},
	{-1, "USD", "-0.01"},
	{-10, "USD", "-0.10"},
	{-100, "USD", "-1.00"},
	{-1000, "USD", "-10.00"},
	{-10000, "USD", "-100.00"},
	{-100000, "USD", "-1,000.00"},
}

func TestDisplayNoSymbol(t *testing.T) {
	for _, v := range TestDisplayNoSymbolData {
		result, err := DisplayNoSymbol(v.Num, v.Alpha)
		if err != nil {
			if err.Error() != v.Output {
				t.Error(err.Error())
			}
		} else if result != v.Output {
			t.Error(result)
		}
	}
}

var TestDisplayWithDecimalData = []struct {
	Num    int
	Alpha  string
	Output string
}{
	{0, "USA", ErrorInvalidISO.Error()},
	{0, "USD", "0.00"},
	{1, "USD", "0.01"},
	{10, "USD", "0.10"},
	{100, "USD", "1.00"},
	{1000, "USD", "10.00"},
	{10000, "USD", "100.00"},
	{100000, "USD", "1000.00"},
	{-1, "USD", "-0.01"},
	{-10, "USD", "-0.10"},
	{-100, "USD", "-1.00"},
	{-1000, "USD", "-10.00"},
	{-10000, "USD", "-100.00"},
	{-100000, "USD", "-1000.00"},
}

func TestDisplayWithDecimal(t *testing.T) {
	for _, v := range TestDisplayWithDecimalData {
		result, err := DisplayWithDecimal(v.Num, v.Alpha)
		if err != nil {
			if err.Error() != v.Output {
				t.Error(err.Error())
			}
		} else if result != v.Output {
			t.Error(result)
		}
	}
}

func TestTopCurrencies(t *testing.T) {
	_, err := TopCurrencies()
	if err != nil {
		t.Error("Could not find all valid currencies from list")
	}
}

var TestListCurrenciesData = []struct {
	Input  []string
	Output interface{}
}{
	{[]string{"USA"}, ErrorInvalidISO},
	{[]string{"USD"}, []Currency{{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}}},
	{[]string{"AED"}, []Currency{{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}}},
	{[]string{"USD", "GBP"}, []Currency{{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, {Unit: "Pound Sterling", Alpha: "GBP", Numeric: "826", Symbol: "£", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}}},
}

func TestListCurrencies(t *testing.T) {
	for _, v := range TestListCurrenciesData {
		result, err := ListCurrencies(v.Input)
		if err != nil {
			if err != v.Output {
				t.Error(err.Error())
			}
		} else if reflect.DeepEqual(result, v.Output) != true {
			t.Error(result)
		}
	}
}

func TestConversions(t *testing.T) {
	var data = []struct {
		Amount int
		Alpha  string
		Output string
	}{
		{0, "USD", "$0.00"},
		{1, "USD", "$0.01"},
		{10, "USD", "$0.10"},
		{100, "USD", "$1.00"},
		{1000, "USD", "$10.00"},
		{10000, "USD", "$100.00"},
		{100000, "USD", "$1,000.00"},
		{0, "CAD", "CA$0.00"},
		{1, "CAD", "CA$0.01"},
		{10, "CAD", "CA$0.10"},
		{100, "CAD", "CA$1.00"},
		{1000, "CAD", "CA$10.00"},
		{10000, "CAD", "CA$100.00"},
		{100000, "CAD", "CA$1,000.00"},
		{0, "MXN", "$0.00"},
		{1, "MXN", "$0.01"},
		{10, "MXN", "$0.10"},
		{100, "MXN", "$1.00"},
		{1000, "MXN", "$10.00"},
		{10000, "MXN", "$100.00"},
		{100000, "MXN", "$1,000.00"},
		{0, "AUD", "$0.00"},
		{1, "AUD", "$0.01"},
		{10, "AUD", "$0.10"},
		{100, "AUD", "$1.00"},
		{1000, "AUD", "$10.00"},
		{10000, "AUD", "$100.00"},
		{100000, "AUD", "$1 000.00"},
		{0, "GBP", "\u00a30.00"},
		{1, "GBP", "\u00a30.01"},
		{10, "GBP", "\u00a30.10"},
		{100, "GBP", "\u00a31.00"},
		{1000, "GBP", "\u00a310.00"},
		{10000, "GBP", "\u00a3100.00"},
		{100000, "GBP", "\u00a31,000.00"},
		{0, "SEK", "\u006b\u00720.00"},
		{1, "SEK", "\u006b\u00720.01"},
		{10, "SEK", "\u006b\u00720.10"},
		{100, "SEK", "\u006b\u00721.00"},
		{1000, "SEK", "\u006b\u007210.00"},
		{10000, "SEK", "\u006b\u0072100.00"},
		{100000, "SEK", "\u006b\u00721,000.00"},
		{0, "PHP", "\u20b10.00"},
		{1, "PHP", "\u20b10.01"},
		{10, "PHP", "\u20b10.10"},
		{100, "PHP", "\u20b11.00"},
		{1000, "PHP", "\u20b110.00"},
		{10000, "PHP", "\u20b1100.00"},
		{100000, "PHP", "\u20b11,000.00"},
		{0, "NOK", "\u006b\u00720.00"},
		{1, "NOK", "\u006b\u00720.01"},
		{10, "NOK", "\u006b\u00720.10"},
		{100, "NOK", "\u006b\u00721.00"},
		{1000, "NOK", "\u006b\u007210.00"},
		{10000, "NOK", "\u006b\u0072100.00"},
		{100000, "NOK", "\u006b\u00721,000.00"},
		{0, "DKK", "\u006b\u0072\u002e0.00"},
		{1, "DKK", "\u006b\u0072\u002e0.01"},
		{10, "DKK", "\u006b\u0072\u002e0.10"},
		{100, "DKK", "\u006b\u0072\u002e1.00"},
		{1000, "DKK", "\u006b\u0072\u002e10.00"},
		{10000, "DKK", "\u006b\u0072\u002e100.00"},
		{100000, "DKK", "\u006b\u0072\u002e1,000.00"},
		{0, "EUR", "\u20ac0.00"},
		{1, "EUR", "\u20ac0.01"},
		{10, "EUR", "\u20ac0.10"},
		{100, "EUR", "\u20ac1.00"},
		{1000, "EUR", "\u20ac10.00"},
		{10000, "EUR", "\u20ac100.00"},
		{100000, "EUR", "\u20ac1,000.00"},
		{0, "JPY", "\u00a50"},
		{1, "JPY", "\u00a51"},
		{10, "JPY", "\u00a510"},
		{100, "JPY", "\u00a5100"},
		{1000, "JPY", "\u00a51000"},
		{10000, "JPY", "\u00a510000"},
		{100000, "JPY", "\u00a5100000"},
	}

	for _, d := range data {
		i, err := GetISOFromAlpha(d.Alpha)
		if err != nil {
			t.Error(err)
		}
		s := FormatCurrency(d.Amount, i)
		if s != d.Output {
			t.Error("not a match")
		}
	}
}
