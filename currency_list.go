package currency

// Currency - struct containing currency variables
type Currency struct {
	Unit      string
	Alpha     string
	Numeric   string
	Symbol    string
	Exponent  int
	Decimal   string
	Separator int
	Delimiter string
}

// CurrencyList - complete list of supported currencies
var CurrencyList = map[string]Currency{
	"USD": {
		Unit:      "US Dollar",
		Alpha:     "USD",
		Numeric:   "840",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Separator: 3,
		Delimiter: ",",
	},
	// "AED": {
	// 	Unit:      "UAE Dirham",
	// 	Alpha:     "AED",
	// 	Numeric:   "784",
	// 	Symbol:    "\u0625\u002E\u062F",
	// 	Exponent:  2,
	// 	Decimal:   ".",
	// 	Separator: 3,
	// 	Delimiter: ",",
	// },
	// "AFN": {
	// 	Unit:      "Afghani",
	// 	Alpha:     "AFN",
	// 	Numeric:   "971",
	// 	Symbol:    "\u060B",
	// 	Exponent:  2,
	// 	Decimal:   ".",
	// 	Separator: 3,
	// 	Delimiter: ",",
	// },
	// "ALL": {
	// 	Unit:      "Lek",
	// 	Alpha:     "ALL",
	// 	Numeric:   "008",
	// 	Symbol:    "\u004c\u0065\u006B",
	// 	Exponent:  2,
	// 	Decimal:   ".",
	// 	Separator: 3,
	// 	Delimiter: ",",
	// },
	// "AMD": {
	// 	Unit:      "Armenian Dram",
	// 	Alpha:     "AMD",
	// 	Numeric:   "051",
	// 	Symbol:    "\u058F",
	// 	Exponent:  2,
	// 	Decimal:   ".",
	// 	Separator: 3,
	// 	Delimiter: ",",
	// },
	// "ANG": {
	// 	Unit:      "Netherlands Antillean Guilder",
	// 	Alpha:     "ANG",
	// 	Numeric:   "532",
	// 	Symbol:    "\u0192",
	// 	Exponent:  2,
	// 	Decimal:   ".",
	// 	Separator: 3,
	// 	Delimiter: ",",
	// },
	// "AOA": {
	// 	Unit:      "Kwanza",
	// 	Alpha:     "AOA",
	// 	Numeric:   "973",
	// 	Symbol:    "Kz",
	// 	Exponent:  2,
	// 	Decimal:   ".",
	// 	Separator: 3,
	// 	Delimiter: ",",
	// },
	// "ARS": {
	// 	Unit:      "Argentine Peso",
	// 	Alpha:     "ARS",
	// 	Numeric:   "032",
	// 	Symbol:    "\u0024",
	// 	Exponent:  2,
	// 	Decimal:   ",",
	// 	Separator: 3,
	// 	Delimiter: ".",
	// },
	// "AUD": {
	// 	Unit:      "Australian Dollar",
	// 	Alpha:     "AUD",
	// 	Numeric:   "036",
	// 	Symbol:    "\u0024",
	// 	Exponent:  2,
	// 	Decimal:   ".",
	// 	Separator: 3,
	// 	Delimiter: " ",
	// },
}
