package countrycontinent

import (
	"strings"
	"testing"
)

// TestCountryGetFullName tests the CountryGetFullName function.
func TestCountryGetFullName(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		expected string
		wantErr  bool
	}{
		{name: "Valid code", code: "US", expected: "United States", wantErr: false},
		{name: "Invalid code", code: "XX", expected: "", wantErr: false},
		{name: "Empty code", code: "", expected: "", wantErr: false},
		{name: "Mixed case code", code: "fr", expected: "France", wantErr: false},
		{name: "Invalid mixed case code", code: "xx", expected: "", wantErr: false},
		{name: "Invalid mixed case code", code: "Xx", expected: "", wantErr: false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CountryGetFullName(tc.code)
			if got != tc.expected {
				t.Errorf("CountryGetFullName(%s) got = %v, want %v", tc.code, got, tc.expected)
			}
		})
	}
}

func TestCountryGetFullNameContinent(t *testing.T) {
	// test cases
	for country, want := range map[string]string{
		"NL": "Netherlands, Europe",
		"US": "United States, North America",
		"MX": "Mexico, North America",
		"PE": "Peru, South America",
		"DM": "Dominica, Caribbean",
		"NZ": "New Zealand (Aotearoa), Oceania",
		"JP": "Japan, Asia",
		"XX": ", ",
	} {
		// expected result
		wantCountry, wantContinent := strings.Split(want, ", ")[0], strings.Split(want, ", ")[1]
		// actual result
		gotCountry, gotContinent := CountryGetFullNameContinent(country)
		// compare expected and actual result
		if gotCountry != wantCountry || gotContinent != wantContinent {
			// report error
			t.Errorf("CountryGetFullNameContinent(%s) = %s, %s; want %s, %s", country, gotCountry, gotContinent, wantCountry, wantContinent)
		}
	}

}

func TestContinentGetCountries(t *testing.T) {
	// Log is only for debugging. To see the output run: go test -v
	t.Logf("Europe %d: %s", len(ContinentGetCountries("Europe")), ContinentGetCountries("Europe"))
	t.Logf("North America %d: %s", len(ContinentGetCountries("North America")), ContinentGetCountries("North America"))
	t.Logf("Central America %d: %s", len(ContinentGetCountries("Central America")), ContinentGetCountries("Central America"))
	t.Logf("South America %d: %s", len(ContinentGetCountries("South America")), ContinentGetCountries("South America"))
	t.Logf("Caribbean %d: %s", len(ContinentGetCountries("Caribbean")), ContinentGetCountries("Caribbean"))
	t.Logf("Oceania %d: %s", len(ContinentGetCountries("Oceania")), ContinentGetCountries("Oceania"))
	t.Logf("Asia %d: %s", len(ContinentGetCountries("Asia")), ContinentGetCountries("Asia"))
	t.Logf("Africa %d: %s", len(ContinentGetCountries("Africa")), ContinentGetCountries("Africa"))
	t.Logf("Antarctica %d: %s", len(ContinentGetCountries("Antarctica")), ContinentGetCountries("Antarctica"))
	t.Logf("XX %d: %s", len(ContinentGetCountries("XX")), ContinentGetCountries("XX"))

	// We want to test if the countries are in the right continent
	for continent, countries := range map[string][]string{
		"Europe":          {"NL", "DE", "FR", "ES", "IT", "GB"},
		"North America":   {"US", "MX"},
		"Central America": {"BZ", "CR", "SV", "GT", "HN", "ni", "PA"},
		"South America":   {"PE"},
		"Caribbean":       {"DM", "CU"},
		"Oceania":         {"NZ"},
		"Asia":            {"JP"},
		"Africa":          {"ZA"},
		"Antarctica":      {"TF"},
		"XX":              {},
	} {
		for _, country := range countries {
			if !StringInSlice(strings.ToUpper(country), ContinentGetCountries(continent)) {
				t.Errorf("ContinentGetCountries(%s) = %s; want %s", continent, ContinentGetCountries(continent), countries)
			}
		}
	}
}

// StringInSlice checks if a string is present in a slice of strings, performing a case-insensitive comparison.
func StringInSlice(country string, s []string) bool {
	uppercaseCountry := strings.ToUpper(country)
	for _, v := range s {
		if strings.ToUpper(v) == uppercaseCountry {
			return true
		}
	}
	return false
}
