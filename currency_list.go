package currency

// Currency - struct containing currency variables
type Currency struct {
	Unit      string
	Alpha     string
	Numeric   string
	Symbol    string
	Exponent  int
	Decimal   string
	Grouping  int
	Delimiter string
}

// CurrencyList - complete list of supported currencies
var CurrencyList = map[string]Currency{
	"AED": {
		Unit:      "UAE Dirham",
		Alpha:     "AED",
		Numeric:   "784",
		Symbol:    "\u0625\u002E\u062F",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"AFN": {
		Unit:      "Afghani",
		Alpha:     "AFN",
		Numeric:   "971",
		Symbol:    "\u060B",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"ALL": {
		Unit:      "Lek",
		Alpha:     "ALL",
		Numeric:   "008",
		Symbol:    "\u004c\u0065\u006B",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"AMD": {
		Unit:      "Armenian Dram",
		Alpha:     "AMD",
		Numeric:   "051",
		Symbol:    "\u058F",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"ANG": {
		Unit:      "Netherlands Antillean Guilder",
		Alpha:     "ANG",
		Numeric:   "532",
		Symbol:    "\u0192",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"AOA": {
		Unit:      "Kwanza",
		Alpha:     "AOA",
		Numeric:   "973",
		Symbol:    "\u004B\u007A",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"ARS": {
		Unit:      "Argentine Peso",
		Alpha:     "ARS",
		Numeric:   "032",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ",",
		Grouping:  3,
		Delimiter: ".",
	},
	"AUD": {
		Unit:      "Australian Dollar",
		Alpha:     "AUD",
		Numeric:   "036",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: " ",
	},
	"AWG": {
		Unit:      "Aruban Florin",
		Alpha:     "AWG",
		Numeric:   "533",
		Symbol:    "\u0192",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"AZN": {
		Unit:      "Azerbaijan Manat",
		Alpha:     "AZN",
		Numeric:   "944",
		Symbol:    "\u20BC",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"BAM": {
		Unit:      "Convertible Mark",
		Alpha:     "BAM",
		Numeric:   "977",
		Symbol:    "\u004B\u004D",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"BBD": {
		Unit:      "Barbados Dollar",
		Alpha:     "BBD",
		Numeric:   "052",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"BDT": {
		Unit:      "Taka",
		Alpha:     "BDT",
		Numeric:   "050",
		Symbol:    "\u0054\u006B",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"BGN": {
		Unit:      "Bulgarian Lev",
		Alpha:     "BGN",
		Numeric:   "975",
		Symbol:    "\u043B\u0432",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"BHD": {
		Unit:      "Bahraini Dinar",
		Alpha:     "BHD",
		Numeric:   "048",
		Symbol:    "\u002e\u062f\u002e\u0628",
		Exponent:  3,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"BIF": {
		Unit:      "Burundi Franc",
		Alpha:     "BIF",
		Numeric:   "108",
		Symbol:    "FBu",
		Exponent:  0,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"BMD": {
		Unit:      "Bermudian Dollar",
		Alpha:     "BMD",
		Numeric:   "060",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"BND": {
		Unit:      "Brunei Dollar",
		Alpha:     "BND",
		Numeric:   "096",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"BOB": {
		Unit:      "Boliviano",
		Alpha:     "BOB",
		Numeric:   "068",
		Symbol:    "$b",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"BRL": {
		Unit:      "Brazilian Real",
		Alpha:     "BRL",
		Numeric:   "986",
		Symbol:    "R$",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"BSD": {
		Unit:      "Bahamian Dollar",
		Alpha:     "BSD",
		Numeric:   "044",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"BTN": {
		Unit:      "Ngultrum",
		Alpha:     "BTN",
		Numeric:   "064",
		Symbol:    "Nu.",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"BWP": {
		Unit:      "Pula",
		Alpha:     "BWP",
		Numeric:   "072",
		Symbol:    "P",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"BYN": {
		Unit:      "Belarusian Ruble",
		Alpha:     "BYN",
		Numeric:   "933",
		Symbol:    "Br",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"BZD": {
		Unit:      "Belize Dollar",
		Alpha:     "BZD",
		Numeric:   "084",
		Symbol:    "BZ$",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"CAD": {
		Unit:      "Canadian Dollar",
		Alpha:     "CAD",
		Numeric:   "124",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"CDF": {
		Unit:      "Congolese Franc",
		Alpha:     "CDF",
		Numeric:   "976",
		Symbol:    "FC",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"CHF": {
		Unit:      "Swiss Franc",
		Alpha:     "CHF",
		Numeric:   "756",
		Symbol:    "CHF",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"CLP": {
		Unit:      "Chilean Peso",
		Alpha:     "CLP",
		Numeric:   "152",
		Symbol:    "\u0024",
		Exponent:  0,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"CNY": {
		Unit:      "Yuan Renminbi",
		Alpha:     "CNY",
		Numeric:   "156",
		Symbol:    "¥",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"COP": {
		Unit:      "Colombian Peso",
		Alpha:     "COP",
		Numeric:   "170",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"CRC": {
		Unit:      "Costa Rican Colon",
		Alpha:     "CRC",
		Numeric:   "188",
		Symbol:    "₡",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"CUC": {
		Unit:      "Peso Convertible",
		Alpha:     "CUC",
		Numeric:   "931",
		Symbol:    "CUC$",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"CUP": {
		Unit:      "Cuban Peso",
		Alpha:     "CUP",
		Numeric:   "192",
		Symbol:    "₱",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"CVE": {
		Unit:      "Cabo Verde Escudo",
		Alpha:     "CVE",
		Numeric:   "132",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"CZK": {
		Unit:      "Czech Koruna",
		Alpha:     "CZK",
		Numeric:   "203",
		Symbol:    "Kč",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"DJF": {
		Unit:      "Djibouti Franc",
		Alpha:     "DJF",
		Numeric:   "262",
		Symbol:    "Fdj",
		Exponent:  0,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"DKK": {
		Unit:      "Danish Krone",
		Alpha:     "DKK",
		Numeric:   "208",
		Symbol:    "kr.",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"DOP": {
		Unit:      "Dominican Peso",
		Alpha:     "DOP",
		Numeric:   "214",
		Symbol:    "RD$",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"DZD": {
		Unit:      "Algerian Dinar",
		Alpha:     "DZD",
		Numeric:   "012",
		Symbol:    "DA",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"EGP": {
		Unit:      "Egyptian Pound",
		Alpha:     "EGP",
		Numeric:   "818",
		Symbol:    "£",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"ERN": {
		Unit:      "Nakfa",
		Alpha:     "ERN",
		Numeric:   "232",
		Symbol:    "Nkf",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"ETB": {
		Unit:      "Ethiopian Birr",
		Alpha:     "ETB",
		Numeric:   "230",
		Symbol:    "Br",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"EUR": {
		Unit:      "Euro",
		Alpha:     "EUR",
		Numeric:   "978",
		Symbol:    "€",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"FJD": {
		Unit:      "Fiji Dollar",
		Alpha:     "FJD",
		Numeric:   "242",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"FKP": {
		Unit:      "Falkland Islands Pound",
		Alpha:     "FKP",
		Numeric:   "238",
		Symbol:    "£",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"GBP": {
		Unit:      "Pound Sterling",
		Alpha:     "GBP",
		Numeric:   "826",
		Symbol:    "£",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"GEL": {
		Unit:      "Lari",
		Alpha:     "GEL",
		Numeric:   "981",
		Symbol:    "₾",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"GHS": {
		Unit:      "Ghana Cedi",
		Alpha:     "GHS",
		Numeric:   "936",
		Symbol:    "GH¢",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"GIP": {
		Unit:      "Gibraltar Pound",
		Alpha:     "GIP",
		Numeric:   "292",
		Symbol:    "£",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"GMD": {
		Unit:      "Dalasi",
		Alpha:     "GMD",
		Numeric:   "270",
		Symbol:    "D",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"GNF": {
		Unit:      "Guinean Franc",
		Alpha:     "GNF",
		Numeric:   "324",
		Symbol:    "FG",
		Exponent:  0,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"GTQ": {
		Unit:      "Quetzal",
		Alpha:     "GTQ",
		Numeric:   "320",
		Symbol:    "Q",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"GYD": {
		Unit:      "Guyana Dollar",
		Alpha:     "GYD",
		Numeric:   "328",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"HKD": {
		Unit:      "Hong Kong Dollar",
		Alpha:     "HKD",
		Numeric:   "344",
		Symbol:    "HK$",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"HNL": {
		Unit:      "Lempira",
		Alpha:     "HNL",
		Numeric:   "340",
		Symbol:    "L",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"HRK": {
		Unit:      "Kuna",
		Alpha:     "HRK",
		Numeric:   "191",
		Symbol:    "kn",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"HTG": {
		Unit:      "Gourde",
		Alpha:     "HTG",
		Numeric:   "332",
		Symbol:    "G",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"HUF": {
		Unit:      "Forint",
		Alpha:     "HUF",
		Numeric:   "348",
		Symbol:    "Ft",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"IDR": {
		Unit:      "Rupiah",
		Alpha:     "IDR",
		Numeric:   "360",
		Symbol:    "Rp",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"ILS": {
		Unit:      "New Israeli Sheqel",
		Alpha:     "ILS",
		Numeric:   "376",
		Symbol:    "₪",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"INR": {
		Unit:      "Indian Rupee",
		Alpha:     "INR",
		Numeric:   "356",
		Symbol:    "₹",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"IQD": {
		Unit:      "Iraqi Dinar",
		Alpha:     "IQD",
		Numeric:   "368",
		Symbol:    "د.ع",
		Exponent:  3,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"IRR": {
		Unit:      "Iranian Rial",
		Alpha:     "IRR",
		Numeric:   "364",
		Symbol:    "﷼",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"ISK": {
		Unit:      "Iceland Krona",
		Alpha:     "ISK",
		Numeric:   "352",
		Symbol:    "kr",
		Exponent:  0,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"JMD": {
		Unit:      "Jamaican Dollar",
		Alpha:     "JMD",
		Numeric:   "388",
		Symbol:    "J$",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"JOD": {
		Unit:      "Jordanian Dinar",
		Alpha:     "JOD",
		Numeric:   "400",
		Symbol:    "د.ا",
		Exponent:  3,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"JPY": {
		Unit:      "Yen",
		Alpha:     "JPY",
		Numeric:   "392",
		Symbol:    "¥",
		Exponent:  0,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"KES": {
		Unit:      "Kenyan Shilling",
		Alpha:     "KES",
		Numeric:   "404",
		Symbol:    "KSh",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"KGS": {
		Unit:      "Som",
		Alpha:     "KGS",
		Numeric:   "417",
		Symbol:    "лв",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"KHR": {
		Unit:      "Riel",
		Alpha:     "KHR",
		Numeric:   "116",
		Symbol:    "៛",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"KMF": {
		Unit:      "Comorian Franc ",
		Alpha:     "KMF",
		Numeric:   "174",
		Symbol:    "CF",
		Exponent:  0,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"KPW": {
		Unit:      "North Korean Won",
		Alpha:     "KPW",
		Numeric:   "408",
		Symbol:    "₩",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"KRW": {
		Unit:      "Won",
		Alpha:     "KRW",
		Numeric:   "410",
		Symbol:    "₩",
		Exponent:  0,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"KWD": {
		Unit:      "Kuwaiti Dinar",
		Alpha:     "KWD",
		Numeric:   "414",
		Symbol:    "ك",
		Exponent:  3,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"KYD": {
		Unit:      "Cayman Islands Dollar",
		Alpha:     "KYD",
		Numeric:   "136",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"KZT": {
		Unit:      "Tenge",
		Alpha:     "KZT",
		Numeric:   "398",
		Symbol:    "₸",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"LAK": {
		Unit:      "Lao Kip",
		Alpha:     "LAK",
		Numeric:   "418",
		Symbol:    "₭",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"LBP": {
		Unit:      "Lebanese Pound",
		Alpha:     "LBP",
		Numeric:   "422",
		Symbol:    "ل.ل",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"LKR": {
		Unit:      "Sri Lanka Rupee",
		Alpha:     "LKR",
		Numeric:   "144",
		Symbol:    "₨",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"LRD": {
		Unit:      "Liberian Dollar",
		Alpha:     "LRD",
		Numeric:   "430",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"LSL": {
		Unit:      "Loti",
		Alpha:     "LSL",
		Numeric:   "426",
		Symbol:    "L",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"LYD": {
		Unit:      "Libyan Dinar",
		Alpha:     "LYD",
		Numeric:   "434",
		Symbol:    "LD",
		Exponent:  3,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"MAD": {
		Unit:      "Moroccan Dirham",
		Alpha:     "MAD",
		Numeric:   "504",
		Symbol:    "د.م.",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"MDL": {
		Unit:      "Moldovan Leu",
		Alpha:     "MDL",
		Numeric:   "498",
		Symbol:    "MDL",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"MGA": {
		Unit:      "Malagasy Ariary",
		Alpha:     "MGA",
		Numeric:   "969",
		Symbol:    "Ar",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"MKD": {
		Unit:      "Denar",
		Alpha:     "MKD",
		Numeric:   "807",
		Symbol:    "ден",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"MMK": {
		Unit:      "Kyat",
		Alpha:     "MMK",
		Numeric:   "104",
		Symbol:    "K",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"MNT": {
		Unit:      "Tugrik",
		Alpha:     "MNT",
		Numeric:   "496",
		Symbol:    "₮",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"MOP": {
		Unit:      "Pataca",
		Alpha:     "MOP",
		Numeric:   "446",
		Symbol:    "MOP$",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"MRU": {
		Unit:      "Ouguiya",
		Alpha:     "MRU",
		Numeric:   "929",
		Symbol:    "UM",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"MUR": {
		Unit:      "Mauritius Rupee",
		Alpha:     "MUR",
		Numeric:   "480",
		Symbol:    "₨",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"MVR": {
		Unit:      "Rufiyaa",
		Alpha:     "MVR",
		Numeric:   "462",
		Symbol:    "Rf",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"MWK": {
		Unit:      "Malawi Kwacha",
		Alpha:     "MWK",
		Numeric:   "454",
		Symbol:    "MK",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"MXN": {
		Unit:      "Mexican Peso",
		Alpha:     "MXN",
		Numeric:   "484",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"MYR": {
		Unit:      "Malaysian Ringgit",
		Alpha:     "MYR",
		Numeric:   "458",
		Symbol:    "RM",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"MZN": {
		Unit:      "Mozambique Metical",
		Alpha:     "MZN",
		Numeric:   "943",
		Symbol:    "MT",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"NAD": {
		Unit:      "Namibia Dollar",
		Alpha:     "NAD",
		Numeric:   "516",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"NGN": {
		Unit:      "Naira",
		Alpha:     "NGN",
		Numeric:   "566",
		Symbol:    "₦",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"NIO": {
		Unit:      "Cordoba Oro",
		Alpha:     "NIO",
		Numeric:   "558",
		Symbol:    "C$",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"NOK": {
		Unit:      "Norwegian Krone",
		Alpha:     "NOK",
		Numeric:   "578",
		Symbol:    "kr",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"NPR": {
		Unit:      "Nepalese Rupee",
		Alpha:     "NPR",
		Numeric:   "524",
		Symbol:    "₨",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"NZD": {
		Unit:      "New Zealand Dollar",
		Alpha:     "NZD",
		Numeric:   "554",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"OMR": {
		Unit:      "Rial Omani",
		Alpha:     "OMR",
		Numeric:   "512",
		Symbol:    "﷼",
		Exponent:  3,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"PAB": {
		Unit:      "Balboa",
		Alpha:     "PAB",
		Numeric:   "590",
		Symbol:    "B/.",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"PEN": {
		Unit:      "Sol",
		Alpha:     "PEN",
		Numeric:   "604",
		Symbol:    "S/.",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"PGK": {
		Unit:      "Kina",
		Alpha:     "PGK",
		Numeric:   "598",
		Symbol:    "K",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"PHP": {
		Unit:      "Philippine Peso",
		Alpha:     "PHP",
		Numeric:   "608",
		Symbol:    "₱",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"PKR": {
		Unit:      "Pakistan Rupee",
		Alpha:     "PKR",
		Numeric:   "586",
		Symbol:    "₨",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"PLN": {
		Unit:      "Zloty",
		Alpha:     "PLN",
		Numeric:   "985",
		Symbol:    "zł",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"PYG": {
		Unit:      "Guarani",
		Alpha:     "PYG",
		Numeric:   "600",
		Symbol:    "Gs",
		Exponent:  0,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"QAR": {
		Unit:      "Qatari Rial",
		Alpha:     "QAR",
		Numeric:   "634",
		Symbol:    "﷼",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"RON": {
		Unit:      "Romanian Leu",
		Alpha:     "RON",
		Numeric:   "946",
		Symbol:    "lei",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"RSD": {
		Unit:      "Serbian Dinar",
		Alpha:     "RSD",
		Numeric:   "941",
		Symbol:    "РСД",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"RUB": {
		Unit:      "Russian Ruble",
		Alpha:     "RUB",
		Numeric:   "643",
		Symbol:    "₽",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"RWF": {
		Unit:      "Rwanda Franc",
		Alpha:     "RWF",
		Numeric:   "646",
		Symbol:    "FRw",
		Exponent:  0,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"SAR": {
		Unit:      "Saudi Riyal",
		Alpha:     "SAR",
		Numeric:   "682",
		Symbol:    "﷼",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"SBD": {
		Unit:      "Solomon Islands Dollar",
		Alpha:     "SBD",
		Numeric:   "090",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"SCR": {
		Unit:      "Seychelles Rupee",
		Alpha:     "SCR",
		Numeric:   "690",
		Symbol:    "₨",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"SDG": {
		Unit:      "Sudanese Pound",
		Alpha:     "SDG",
		Numeric:   "938",
		Symbol:    "ج.س.",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"SEK": {
		Unit:      "Swedish Krona",
		Alpha:     "SEK",
		Numeric:   "752",
		Symbol:    "kr",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"SGD": {
		Unit:      "Singapore Dollar",
		Alpha:     "SGD",
		Numeric:   "702",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"SHP": {
		Unit:      "Saint Helena Pound",
		Alpha:     "SHP",
		Numeric:   "654",
		Symbol:    "£",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"SLL": {
		Unit:      "Leone",
		Alpha:     "SLL",
		Numeric:   "694",
		Symbol:    "Le",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"SOS": {
		Unit:      "Somali Shilling",
		Alpha:     "SOS",
		Numeric:   "706",
		Symbol:    "S",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"SRD": {
		Unit:      "Surinam Dollar",
		Alpha:     "SRD",
		Numeric:   "968",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"SSP": {
		Unit:      "South Sudanese Pound",
		Alpha:     "SSP",
		Numeric:   "728",
		Symbol:    "£",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"STN": {
		Unit:      "Dobra",
		Alpha:     "STN",
		Numeric:   "930",
		Symbol:    "Db",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"SVC": {
		Unit:      "El Salvador Colon",
		Alpha:     "SVC",
		Numeric:   "222",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"SYP": {
		Unit:      "Syrian Pound",
		Alpha:     "SYP",
		Numeric:   "760",
		Symbol:    "£",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"SZL": {
		Unit:      "Lilangeni",
		Alpha:     "SZL",
		Numeric:   "748",
		Symbol:    "L",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"THB": {
		Unit:      "Baht",
		Alpha:     "THB",
		Numeric:   "764",
		Symbol:    "฿",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"TJS": {
		Unit:      "Somoni",
		Alpha:     "TJS",
		Numeric:   "972",
		Symbol:    "ЅM",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"TMT": {
		Unit:      "Turkmenistan New Manat",
		Alpha:     "TMT",
		Numeric:   "934",
		Symbol:    "T",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"TND": {
		Unit:      "Tunisian Dinar",
		Alpha:     "TND",
		Numeric:   "788",
		Symbol:    "د.ت",
		Exponent:  3,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"TOP": {
		Unit:      "Pa’anga",
		Alpha:     "TOP",
		Numeric:   "776",
		Symbol:    "T$",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"TRY": {
		Unit:      "Turkish Lira",
		Alpha:     "TRY",
		Numeric:   "949",
		Symbol:    "₺",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"TTD": {
		Unit:      "Trinidad and Tobago Dollar",
		Alpha:     "TTD",
		Numeric:   "780",
		Symbol:    "TT$",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"TWD": {
		Unit:      "New Taiwan Dollar",
		Alpha:     "TWD",
		Numeric:   "901",
		Symbol:    "NT$",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"TZS": {
		Unit:      "Tanzanian Shilling",
		Alpha:     "TZS",
		Numeric:   "834",
		Symbol:    "TSh",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"UAH": {
		Unit:      "Hryvnia",
		Alpha:     "UAH",
		Numeric:   "980",
		Symbol:    "₴",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"UGX": {
		Unit:      "Uganda Shilling",
		Alpha:     "UGX",
		Numeric:   "800",
		Symbol:    "UGX",
		Exponent:  0,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"USD": {
		Unit:      "US Dollar",
		Alpha:     "USD",
		Numeric:   "840",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"UYU": {
		Unit:      "Peso Uruguayo",
		Alpha:     "UYU",
		Numeric:   "858",
		Symbol:    "$U",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"UZS": {
		Unit:      "Uzbekistan Sum",
		Alpha:     "UZS",
		Numeric:   "860",
		Symbol:    "лв",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"VES": {
		Unit:      "Bolívar Soberano",
		Alpha:     "VES",
		Numeric:   "928",
		Symbol:    "Bs. S",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"VND": {
		Unit:      "Dong",
		Alpha:     "VND",
		Numeric:   "704",
		Symbol:    "₫",
		Exponent:  0,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"VUV": {
		Unit:      "Vatu",
		Alpha:     "VUV",
		Numeric:   "548",
		Symbol:    "VT",
		Exponent:  0,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"WST": {
		Unit:      "Tala",
		Alpha:     "WST",
		Numeric:   "882",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"XAF": {
		Unit:      "CFA Franc BEAC",
		Alpha:     "XAF",
		Numeric:   "950",
		Symbol:    "FCFA",
		Exponent:  0,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"XCD": {
		Unit:      "East Caribbean Dollar",
		Alpha:     "XCD",
		Numeric:   "951",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"XOF": {
		Unit:      "CFA Franc BCEAO",
		Alpha:     "XOF",
		Numeric:   "952",
		Symbol:    "CFA",
		Exponent:  0,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"XPF": {
		Unit:      "CFP Franc",
		Alpha:     "XPF",
		Numeric:   "953",
		Symbol:    "₣",
		Exponent:  0,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"YER": {
		Unit:      "Yemeni Rial",
		Alpha:     "YER",
		Numeric:   "886",
		Symbol:    "﷼",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"ZAR": {
		Unit:      "Rand",
		Alpha:     "ZAR",
		Numeric:   "710",
		Symbol:    "R",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"ZMW": {
		Unit:      "Zambian Kwacha",
		Alpha:     "ZMW",
		Numeric:   "967",
		Symbol:    "ZK",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
	"ZWL": {
		Unit:      "Zimbabwe Dollar",
		Alpha:     "ZWL",
		Numeric:   "932",
		Symbol:    "\u0024",
		Exponent:  2,
		Decimal:   ".",
		Grouping:  3,
		Delimiter: ",",
	},
}
