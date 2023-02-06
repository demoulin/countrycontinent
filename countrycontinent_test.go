package countrycontinent

import (
	"testing"
)

func TestCountryGetFullName(t *testing.T) {
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
		if CountryGetFullName(country) != fullname {
			t.Errorf("CountryGetFullName(%s) = %s; want %s", country, CountryGetFullName(country), fullname)
		}
	}
}

func TestCountryGetContinent(t *testing.T) {
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
		if CountryGetContinent(country) != continent {
			t.Errorf("CountryGetContinent(%s) = %s; want %s", country, CountryGetContinent(country), continent)
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

func StringInSlice(country string, s []string) bool {
	for _, v := range s {
		if v == country {
			return true
		}
	}
	return false
}
