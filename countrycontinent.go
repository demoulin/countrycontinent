// Package countrycontinent provides functionality to retrieve information about countries and continents.
// It allows mapping country codes to country names and continents, as well as retrieving countries by continent.
//
// Usage:
//
//   - CountryGetFullName(countryCode string) string
//     Returns the full name of a country given its country code.
//
//   - CountryGetFullNameContinent(countryCode string) (string, string)
//     Returns the full name and continent of a country given its country code.
//
//   - CountryGetContinent(countryCode string) string
//     Returns the continent of a country given its country code.
//
//   - ContinentGetCountries(continent string) []string
//     Returns a list of country codes belonging to a given continent.
package countrycontinent

import (
	"strings"
)

// CountryContinent is a struct that holds the country code, country name and continent
type CountryContinent struct {
	CountryCode string // ISO 3166-1 alpha-2 country code
	CountryName string // Full name of the country
	Continent   string // Continent to which the country belongs
}

// countryContinent is a slice of CountryContinent
var countryContinent = []CountryContinent{
	{"AD", "Andorra", "Europe"},
	{"AE", "United Arab Emirates", "Asia"},
	{"AF", "Afghanistan", "Asia"},
	{"AG", "Antigua and Barbuda", "Caribbean"},
	{"AI", "Anguilla", "Caribbean"},
	{"AL", "Albania", "Europe"},
	{"AM", "Armenia", "Asia"},
	{"AN", "Netherlands Antilles", "Caribbean"},
	{"AO", "Angola", "Africa"},
	{"AR", "Argentina", "South America"},
	{"AS", "American Samoa", "Oceania"},
	{"AT", "Austria", "Europe"},
	{"AU", "Australia", "Oceania"},
	{"AW", "Aruba", "Caribbean"},
	{"AZ", "Azerbaijan", "Asia"},
	{"BA", "Bosnia and Herzegovina", "Europe"},
	{"BB", "Barbados", "Caribbean"},
	{"BD", "Bangladesh", "Asia"},
	{"BE", "Belgium", "Europe"},
	{"BF", "Burkina Faso", "Africa"},
	{"BG", "Bulgaria", "Europe"},
	{"BH", "Bahrain", "Asia"},
	{"BI", "Burundi", "Africa"},
	{"BJ", "Benin", "Africa"},
	{"BM", "Bermuda", "Caribbean"},
	{"BN", "Brunei Darussalam", "Asia"},
	{"BO", "Bolivia", "South America"},
	{"BR", "Brazil", "South America"},
	{"BS", "Bahamas", "Caribbean"},
	{"BT", "Bhutan", "Asia"},
	{"BW", "Botswana", "Africa"},
	{"BY", "Belarus", "Europe"},
	{"BZ", "Belize", "Central America"},
	{"CA", "Canada", "North America"},
	{"CC", "Cocos (Keeling) Islands", "Asia"},
	{"CF", "Central African Republic", "Africa"},
	{"CG", "Congo", "Africa"},
	{"CH", "Switzerland", "Europe"},
	{"CI", "Cote D'Ivoire (Ivory Coast)", "Africa"},
	{"CK", "Cook Islands", "Oceania"},
	{"CL", "Chile", "South America"},
	{"CM", "Cameroon", "Africa"},
	{"CN", "China", "Asia"},
	{"CO", "Colombia", "South America"},
	{"CR", "Costa Rica", "Central America"},
	{"CU", "Cuba", "Caribbean"},
	{"CV", "Cape Verde", "Africa"},
	{"CX", "Christmas Island", "Asia"},
	{"CY", "Cyprus", "Asia"},
	{"CZ", "Czech Republic", "Europe"},
	{"DE", "Germany", "Europe"},
	{"DJ", "Djibouti", "Africa"},
	{"DK", "Denmark", "Europe"},
	{"DM", "Dominica", "Caribbean"},
	{"DO", "Dominican Republic", "Caribbean"},
	{"DZ", "Algeria", "Africa"},
	{"EC", "Ecuador", "South America"},
	{"EE", "Estonia", "Europe"},
	{"EG", "Egypt", "Africa"},
	{"EH", "Western Sahara", "Africa"},
	{"ER", "Eritrea", "Africa"},
	{"ES", "Spain", "Europe"},
	{"ET", "Ethiopia", "Africa"},
	{"FI", "Finland", "Europe"},
	{"FJ", "Fiji", "Oceania"},
	{"FK", "Falkland Islands (Malvinas)", "South America"},
	{"FM", "Micronesia", "Oceania"},
	{"FO", "Faroe Islands", "Europe"},
	{"FR", "France", "Europe"},
	{"GA", "Gabon", "Africa"},
	{"GB", "Great Britain (UK)", "Europe"},
	{"GD", "Grenada", "Caribbean"},
	{"GE", "Georgia", "Asia"},
	{"GF", "French Guiana", "South America"},
	{"GH", "Ghana", "Africa"},
	{"GI", "Gibraltar", "Europe"},
	{"GL", "Greenland", "North America"},
	{"GM", "Gambia", "Africa"},
	{"GN", "Guinea", "Africa"},
	{"GP", "Guadeloupe", "Caribbean"},
	{"GQ", "Equatorial Guinea", "Africa"},
	{"GR", "Greece", "Europe"},
	{"GS", "S. Georgia and S. Sandwich Isls.", "South America"},
	{"GT", "Guatemala", "Central America"},
	{"GU", "Guam", "Oceania"},
	{"GW", "Guinea-Bissau", "Africa"},
	{"GY", "Guyana", "South America"},
	{"HK", "Hong Kong", "Asia"},
	{"HN", "Honduras", "Central America"},
	{"HR", "Croatia (Hrvatska)", "Europe"},
	{"HT", "Haiti", "Caribbean"},
	{"HU", "Hungary", "Europe"},
	{"ID", "Indonesia", "Asia"},
	{"IE", "Ireland", "Europe"},
	{"IL", "Israel", "Asia"},
	{"IN", "India", "Asia"},
	{"IO", "British Indian Ocean Territory", "Asia"},
	{"IQ", "Iraq", "Asia"},
	{"IR", "Iran", "Asia"},
	{"IS", "Iceland", "Europe"},
	{"IT", "Italy", "Europe"},
	{"JM", "Jamaica", "Caribbean"},
	{"JO", "Jordan", "Asia"},
	{"JP", "Japan", "Asia"},
	{"KE", "Kenya", "Africa"},
	{"KG", "Kyrgyzstan", "Asia"},
	{"KH", "Cambodia", "Asia"},
	{"KI", "Kiribati", "Oceania"},
	{"KM", "Comoros", "Africa"},
	{"KN", "Saint Kitts and Nevis", "Caribbean"},
	{"KP", "Korea (North)", "Asia"},
	{"KR", "Korea (South)", "Asia"},
	{"KW", "Kuwait", "Asia"},
	{"KY", "Cayman Islands", "Caribbean"},
	{"KZ", "Kazakhstan", "Asia"},
	{"LA", "Laos", "Asia"},
	{"LB", "Lebanon", "Asia"},
	{"LC", "Saint Lucia", "Caribbean"},
	{"LI", "Liechtenstein", "Europe"},
	{"LK", "Sri Lanka", "Asia"},
	{"LR", "Liberia", "Africa"},
	{"LS", "Lesotho", "Africa"},
	{"LT", "Lithuania", "Europe"},
	{"LU", "Luxembourg", "Europe"},
	{"LV", "Latvia", "Europe"},
	{"LY", "Libya", "Africa"},
	{"MA", "Morocco", "Africa"},
	{"MC", "Monaco", "Europe"},
	{"MD", "Moldova", "Europe"},
	{"MG", "Madagascar", "Africa"},
	{"MH", "Marshall Islands", "Oceania"},
	{"MK", "Macedonia", "Europe"},
	{"ML", "Mali", "Africa"},
	{"MM", "Myanmar", "Asia"},
	{"MN", "Mongolia", "Asia"},
	{"MO", "Macau", "Asia"},
	{"MP", "Northern Mariana Islands", "Oceania"},
	{"MQ", "Martinique", "Caribbean"},
	{"MR", "Mauritania", "Africa"},
	{"MS", "Montserrat", "Caribbean"},
	{"MT", "Malta", "Europe"},
	{"MU", "Mauritius", "Africa"},
	{"MV", "Maldives", "Asia"},
	{"MW", "Malawi", "Africa"},
	{"MX", "Mexico", "North America"},
	{"MY", "Malaysia", "Asia"},
	{"MZ", "Mozambique", "Africa"},
	{"NA", "Namibia", "Africa"},
	{"NC", "New Caledonia", "Oceania"},
	{"NE", "Niger", "Africa"},
	{"NF", "Norfolk Island", "Oceania"},
	{"NG", "Nigeria", "Africa"},
	{"NI", "Nicaragua", "Central America"},
	{"NL", "Netherlands", "Europe"},
	{"NO", "Norway", "Europe"},
	{"NP", "Nepal", "Asia"},
	{"NR", "Nauru", "Oceania"},
	{"NU", "Niue", "Oceania"},
	{"NZ", "New Zealand (Aotearoa)", "Oceania"},
	{"OM", "Oman", "Asia"},
	{"PA", "Panama", "Central America"},
	{"PE", "Peru", "South America"},
	{"PF", "French Polynesia", "Oceania"},
	{"PG", "Papua New Guinea", "Oceania"},
	{"PH", "Philippines", "Asia"},
	{"PK", "Pakistan", "Asia"},
	{"PL", "Poland", "Europe"},
	{"PM", "St. Pierre and Miquelon", "North America"},
	{"PN", "Pitcairn", "Oceania"},
	{"PR", "Puerto Rico", "Caribbean"},
	{"PT", "Portugal", "Europe"},
	{"PW", "Palau", "Oceania"},
	{"PY", "Paraguay", "South America"},
	{"QA", "Qatar", "Asia"},
	{"RE", "Reunion", "Africa"},
	{"RO", "Romania", "Europe"},
	{"RU", "Russian Federation", "Europe"},
	{"RW", "Rwanda", "Africa"},
	{"SA", "Saudi Arabia", "Asia"},
	{"SB", "Solomon Islands", "Oceania"},
	{"SC", "Seychelles", "Africa"},
	{"SD", "Sudan", "Africa"},
	{"SE", "Sweden", "Europe"},
	{"SG", "Singapore", "Asia"},
	{"SH", "St. Helena", "Africa"},
	{"SI", "Slovenia", "Europe"},
	{"SJ", "Svalbard and Jan Mayen Islands", "Europe"},
	{"SK", "Slovakia", "Europe"},
	{"SL", "Sierra Leone", "Africa"},
	{"SM", "San Marino", "Europe"},
	{"SN", "Senegal", "Africa"},
	{"SO", "Somalia", "Africa"},
	{"SR", "Suriname", "South America"},
	{"ST", "Sao Tome and Principe", "Africa"},
	{"SV", "El Salvador", "Central America"},
	{"SY", "Syrian Arab Republic", "Asia"},
	{"SZ", "Swaziland", "Africa"},
	{"TC", "Turks and Caicos Islands", "Caribbean"},
	{"TD", "Chad", "Africa"},
	{"TF", "French Southern Territories", "Antarctica"},
	{"TG", "Togo", "Africa"},
	{"TH", "Thailand", "Asia"},
	{"TJ", "Tajikistan", "Asia"},
	{"TK", "Tokelau", "Oceania"},
	{"TM", "Turkmenistan", "Asia"},
	{"TN", "Tunisia", "Africa"},
	{"TO", "Tonga", "Oceania"},
	{"TP", "East Timor", "Asia"},
	{"TR", "Turkey", "Asia"},
	{"TT", "Trinidad and Tobago", "Caribbean"},
	{"TV", "Tuvalu", "Oceania"},
	{"TW", "Taiwan", "Asia"},
	{"TZ", "Tanzania", "Africa"},
	{"UA", "Ukraine", "Europe"},
	{"UG", "Uganda", "Africa"},
	{"UM", "United States Minor Outlying Islands", "Oceania"},
	{"US", "United States", "North America"},
	{"UY", "Uruguay", "South America"},
	{"UZ", "Uzbekistan", "Asia"},
	{"VA", "Vatican City State (Holy See)", "Europe"},
	{"VC", "St. Vincent and the Grenadines", "Caribbean"},
	{"VE", "Venezuela", "South America"},
	{"VG", "Virgin Islands (British)", "Caribbean"},
	{"VI", "Virgin Islands (U.S.)", "Caribbean"},
	{"VN", "Viet Nam", "Asia"},
	{"VU", "Vanuatu", "Oceania"},
	{"WF", "Wallis and Futuna Islands", "Oceania"},
	{"WS", "Samoa", "Oceania"},
	{"YE", "Yemen", "Asia"},
	{"YT", "Mayotte", "Africa"},
	{"YU", "Yugoslavia", "Europe"},
	{"ZA", "South Africa", "Africa"},
	{"ZM", "Zambia", "Africa"},
	{"ZR", "Zaire", "Africa"},
	{"ZW", "Zimbabwe", "Africa"},
}

var countryMap map[string]CountryContinent

func init() {
	countryMap = make(map[string]CountryContinent)
	for _, country := range countryContinent {
		countryMap[country.CountryCode] = country
	}
}

// CountryGetFullName returns the full name of the country with the given country code.
// This function will return an empty string if the country code is not found.
func CountryGetFullName(countryCode string) string {
	countryCode = strings.ToUpper(countryCode)
	if country, ok := countryMap[countryCode]; ok {
		return country.CountryName
	}
	return ""
}

// CountryGetFullNameContinent returns the full name and continent of the country with the given country code.
func CountryGetFullNameContinent(countryCode string) (string, string) {
	countryCode = strings.ToUpper(countryCode)
	if country, ok := countryMap[countryCode]; ok {
		return country.CountryName, country.Continent
	}
	return "", ""
}

// CountryGetContinent returns the continent of a country from its country code (case sensitive)
// This function will return an empty string if the country code is not found.
func CountryGetContinent(countryCode string) string {
	countryCode = strings.ToUpper(countryCode)
	if country, ok := countryMap[countryCode]; ok {
		return country.Continent
	}
	return ""
}

// ContinentGetCountries returns a list of countries in a continent from its continent name (case sensitive)
func ContinentGetCountries(continent string) (countries []string) {
	for _, country := range countryContinent {
		if country.Continent == continent {
			countries = append(countries, country.CountryCode)
		}
	}
	return countries
}
