package countrycontinent

import (
	"testing"
)

// TestCountryGetFullName tests the CountryGetFullName function.
func TestCountryGetFullName(t *testing.T) {
	// Test cases are represented as a map of country codes to full names.
	for country, fullname := range map[string]string{
		"NL": "Netherlands",
		"US": "United States",
		"MX": "Mexico",
		"PE": "Peru",
		"DM": "Dominica",
		"NZ": "New Zealand (Aotearoa)",
		"JP": "Japan",
		"XX": "",
	} {
		// The test function checks that the country code returns the expected full name.
		if CountryGetFullName(country) != fullname {
			// If the expected and actual values are different, the test fails.
			t.Errorf("CountryGetFullName(%s) = %s; want %s", country, CountryGetFullName(country), fullname)
		}
	}
}

func TestCountryGetContinent(t *testing.T) {
	// test cases
	for country, continent := range map[string]string{
		"NL": "Europe",
		"US": "North America",
		"MX": "North America",
		"PE": "South America",
		"DM": "Caribbean",
		"NZ": "Oceania",
		"JP": "Asia",
		"XX": "",
	} {
		// expected result
		want := continent
		// actual result
		got := CountryGetContinent(country)
		// compare expected and actual result
		if got != want {
			// report error
			t.Errorf("CountryGetContinent(%s) = %s; want %s", country, got, want)
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
		"Europe":        {"NL", "DE", "FR", "ES", "IT", "GB"},
		"North America": {"US", "MX"},
		"South America": {"PE"},
		"Caribbean":     {"DM", "CU"},
		"Oceania":       {"NZ"},
		"Asia":          {"JP"},
		"XX":            {},
	} {
		for _, country := range countries {
			if !StringInSlice(country, ContinentGetCountries(continent)) {
				t.Errorf("ContinentGetCountries(%s) = %s; want %s", continent, ContinentGetCountries(continent), countries)
			}
		}
	}
}

// StringInSlice checks if a string is present in a slice of strings
func StringInSlice(country string, s []string) bool {
	// Iterate through the slice
	for _, v := range s {
		// Check if the country is in the slice
		if v == country {
			return true
		}
	}
	// If the country is not in the slice, return false
	return false
}
