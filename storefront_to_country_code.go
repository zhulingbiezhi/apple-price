package apple_price

import (
	"strings"
)

var storefronts = []string{
	"SUR", "TTO", "TUN", "XKS", "NGA", "COG", "BRB", "BFA", "AGO", "ARE", "ARG", "ARM", "AUS", "AUT", "AZE", "BEL",
	"BEN", "BGR", "BHR", "BHS", "BIH", "BLR", "BLZ", "BOL", "BRA", "BRN", "BWA", "CAN", "CHE", "CHL", "CHN", "CIV",
	"CMR", "COD", "COL", "CRI", "CYP", "CZE", "DEU", "DNK", "DOM", "DZA", "ECU", "EGY", "ESP", "EST", "FIN", "FJI",
	"FRA", "GAB", "GBR", "GEO", "GRC", "GTM", "GUY", "HKG", "HND", "HRV", "HUN", "IDN", "IND", "IRL", "IRQ", "ISL",
	"ISR", "ITA", "JAM", "JOR", "JPN", "KAZ", "KEN", "KGZ", "KHM", "KOR", "KWT", "LAO", "LBN", "LBR", "LBY", "LCA",
	"LKA", "LTU", "LUX", "LVA", "MAC", "MAR", "MDA", "MDV", "MEX", "MKD", "MLI", "MLT", "MMR", "MNE", "MNG", "MOZ",
	"MRT", "MUS", "MYS", "NAM", "NIC", "NLD", "NOR", "NPL", "NZL", "OMN", "PAK", "PAN", "PER", "PHL", "PLW", "POL",
	"PRT", "PRY", "QAT", "ROU", "RUS", "RWA", "SAU", "SEN", "SGP", "SLV", "SRB", "SVK", "SVN", "SWE", "THA", "TJK",
	"TON", "TUR", "TWN", "TZA", "UGA", "UKR", "URY", "USA", "UZB", "VNM", "YEM", "ZAF", "ZMB", "ZWE",
}

// StorefrontToCountryCodeMap Storefront to Country Code mapping with 3-character identifiers
var StorefrontToCountryCodeMap = map[string]string{
	"ZAF": "ZA", // South Africa
	"PRT": "PT", // Portugal
	"GBR": "GB", // United Kingdom
	"ITA": "IT", // Italy
	"CIV": "CI", // Côte d'Ivoire
	"NPL": "NP", // Nepal
	"ARE": "AE", // United Arab Emirates
	"BHS": "BS", // Bahamas
	"CYP": "CY", // Cyprus
	"PAK": "PK", // Pakistan
	"JPN": "JP", // Japan
	"SWE": "SE", // Sweden
	"VNM": "VN", // Vietnam
	"SAU": "SA", // Saudi Arabia
	"PER": "PE", // Peru
	"ROU": "RO", // Romania
	"TWN": "TW", // Taiwan
	"PAN": "PA", // Panama
	"BLZ": "BZ", // Belize
	"UGA": "UG", // Uganda
	"MEX": "MX", // Mexico
	"MDV": "MV", // Maldives
	"ZMB": "ZM", // Zambia
	"HKG": "HK", // Hong Kong
	"CHE": "CH", // Switzerland
	"LAO": "LA", // Laos
	"DZA": "DZ", // Algeria
	"KOR": "KR", // South Korea
	"CAN": "CA", // Canada
	"RUS": "RU", // Russia
	"CRI": "CR", // Costa Rica
	"SRB": "RS", // Serbia
	"LVA": "LV", // Latvia
	"NOR": "NO", // Norway
	"LKA": "LK", // Sri Lanka
	"MAR": "MA", // Morocco
	"IRL": "IE", // Ireland
	"CZE": "CZ", // Czech Republic
	"HND": "HN", // Honduras
	"BEL": "BE", // Belgium
	"SLV": "SV", // El Salvador
	"JOR": "JO", // Jordan
	"BEN": "BJ", // Benin
	"PRY": "PY", // Paraguay
	"SVK": "SK", // Slovakia
	"LTU": "LT", // Lithuania
	"BRN": "BN", // Brunei
	"IND": "IN", // India
	"BWA": "BW", // Botswana
	"BGR": "BG", // Bulgaria
	"KWT": "KW", // Kuwait
	"TUR": "TR", // Turkey
	"ARG": "AR", // Argentina
	"NLD": "NL", // Netherlands
	"POL": "PL", // Poland
	"NIC": "NI", // Nicaragua
	"USA": "US", // United States
	"CMR": "CM", // Cameroon
	"SVN": "SI", // Slovenia
	"IRQ": "IQ", // Iraq
	"THA": "TH", // Thailand
	"BHR": "BH", // Bahrain
	"BIH": "BA", // Bosnia and Herzegovina
	"PHL": "PH", // Philippines
	"LBN": "LB", // Lebanon
	"KEN": "KE", // Kenya
	"IDN": "ID", // Indonesia
	"DOM": "DO", // Dominican Republic
	"GUY": "GY", // Guyana
	"BLR": "BY", // Belarus
	"NAM": "NA", // Namibia
	"AGO": "AO", // Angola
	"HUN": "HU", // Hungary
	"UZB": "UZ", // Uzbekistan
	"ISL": "IS", // Iceland
	"MAC": "MO", // Macau
	"BRA": "BR", // Brazil
	"ISR": "IL", // Israel
	"JAM": "JM", // Jamaica
	"PLW": "PW", // Palau
	"SEN": "SN", // Senegal
	"QAT": "QA", // Qatar
	"ZWE": "ZW", // Zimbabwe
	"GAB": "GA", // Gabon
	"MMR": "MM", // Myanmar
	"AUS": "AU", // Australia
	"KHM": "KH", // Cambodia
	"ECU": "EC", // Ecuador
	"MYS": "MY", // Malaysia
	"TZA": "TZ", // Tanzania
	"ESP": "ES", // Spain
	"KAZ": "KZ", // Kazakhstan
	"ARM": "AM", // Armenia
	"LUX": "LU", // Luxembourg
	"URY": "UY", // Uruguay
	"MUS": "MU", // Mauritius
	"DEU": "DE", // Germany
	"GRC": "GR", // Greece
	"FRA": "FR", // France
	"MRT": "MR", // Mauritania
	"FIN": "FI", // Finland
	"SGP": "SG", // Singapore
	"AUT": "AT", // Austria
	"LCA": "LC", // Saint Lucia
	"MNE": "ME", // Montenegro
	"GTM": "GT", // Guatemala
	"EST": "EE", // Estonia
	"TON": "TO", // Tonga
	"AZE": "AZ", // Azerbaijan
	"OMN": "OM", // Oman
	"GEO": "GE", // Georgia
	"RWA": "RW", // Rwanda
	"COD": "CD", // Congo
	"TJK": "TJ", // Tajikistan
	"KGZ": "KG", // Kyrgyzstan
	"HRV": "HR", // Croatia
	"EGY": "EG", // Egypt
	"MLI": "ML", // Mali
	"DNK": "DK", // Denmark
	"YEM": "YE", // Yemen
	"MNG": "MN", // Mongolia
	"COL": "CO", // Colombia
	"CHL": "CL", // Chile
	"NZL": "NZ", // New Zealand
	"CHN": "CN", // China
	"UKR": "UA", // Ukraine
	"LBR": "LR", // Liberia
	"BOL": "BO", // Bolivia
	"FJI": "FJ", // Fiji
	"LBY": "LY", // Libya
	"MDA": "MD", // Moldova
	"MLT": "MT", // Malta
	"MOZ": "MZ", // Mozambique
	"MKD": "MK", // North Macedonia
	"ALB": "AL", // Albania
	"BFA": "BF", // Burkina Faso
	"BRB": "BB", // Barbados
	"COG": "CG", // Congo
	"NGA": "NG", // Nigeria
	"SUR": "SR", // Suriname
	"TTO": "TT", // Trinidad and Tobago
	"TUN": "TN", // Tunisia
	"XKS": "XK", // Kosovo
}

// ConvertStorefrontToCountryCode converts a 3-character Apple storefront identifier to a country code.
func ConvertStorefrontToCountryCode(storefront string) string {
	// Convert storefront to uppercase for consistency
	storefront = strings.ToUpper(storefront)

	// Check if the storefront identifier exists in the map
	code, exists := StorefrontToCountryCodeMap[storefront]
	if !exists {
		return ConvertResultUnknown
	}
	return code
}
