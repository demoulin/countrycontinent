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
		{name: "Invalid code", code: "XX", expected: "", wantErr: true},
		{name: "Empty code", code: "", expected: "", wantErr: true},
		{name: "Mixed case code", code: "fr", expected: "France", wantErr: false},
		{name: "Invalid mixed case code", code: "xx", expected: "", wantErr: true},
		{name: "Invalid mixed case code", code: "Xx", expected: "", wantErr: true},
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
	tests := []struct {
		name          string
		code          string
		wantCountry   string
		wantContinent string
		wantErr       bool
	}{
		{name: "Valid code", code: "NL", wantCountry: "Netherlands", wantContinent: "Europe", wantErr: false},
		{name: "Valid code", code: "US", wantCountry: "United States", wantContinent: "North America", wantErr: false},
		{name: "Valid code", code: "MX", wantCountry: "Mexico", wantContinent: "North America", wantErr: false},
		{name: "Valid code", code: "PE", wantCountry: "Peru", wantContinent: "South America", wantErr: false},
		{name: "Valid code", code: "DM", wantCountry: "Dominica", wantContinent: "Caribbean", wantErr: false},
		{name: "Valid code", code: "NZ", wantCountry: "New Zealand (Aotearoa)", wantContinent: "Oceania", wantErr: false},
		{name: "Valid code", code: "JP", wantCountry: "Japan", wantContinent: "Asia", wantErr: false},
		{name: "Invalid code", code: "XX", wantCountry: "", wantContinent: "", wantErr: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotCountry, gotContinent := CountryGetFullNameContinent(tc.code)
			if gotCountry != tc.wantCountry || gotContinent != tc.wantContinent {
				t.Errorf("CountryGetFullNameContinent(%s) = %s, %s; want %s, %s", tc.code, gotCountry, gotContinent, tc.wantCountry, tc.wantContinent)
			}
		})
	}
}

func TestContinentGetCountries(t *testing.T) {
	tests := []struct {
		name      string
		continent string
		countries []string
	}{
		{name: "Europe", continent: "Europe", countries: []string{"NL", "DE", "FR", "ES", "IT", "GB"}},
		{name: "North America", continent: "North America", countries: []string{"US", "MX"}},
		{name: "Central America", continent: "Central America", countries: []string{"BZ", "CR", "SV", "GT", "HN", "NI", "PA"}},
		{name: "South America", continent: "South America", countries: []string{"PE"}},
		{name: "Caribbean", continent: "Caribbean", countries: []string{"DM", "CU"}},
		{name: "Oceania", continent: "Oceania", countries: []string{"NZ"}},
		{name: "Asia", continent: "Asia", countries: []string{"JP"}},
		{name: "Africa", continent: "Africa", countries: []string{"ZA"}},
		{name: "Antarctica", continent: "Antarctica", countries: []string{"TF"}},
		{name: "Invalid continent", continent: "XX", countries: []string{}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ContinentGetCountries(tc.continent)
			for _, country := range tc.countries {
				if !StringInSlice(strings.ToUpper(country), got) {
					t.Errorf("ContinentGetCountries(%s) = %s; want %s", tc.continent, got, tc.countries)
				}
			}
		})
	}
}

func TestCountryGetContinent(t *testing.T) {
	tests := []struct {
		name    string
		code    string
		want    string
		wantErr bool
	}{
		{name: "Valid code", code: "NL", want: "Europe", wantErr: false},
		{name: "Valid code", code: "US", want: "North America", wantErr: false},
		{name: "Valid code", code: "MX", want: "North America", wantErr: false},
		{name: "Valid code", code: "PE", want: "South America", wantErr: false},
		{name: "Valid code", code: "DM", want: "Caribbean", wantErr: false},
		{name: "Valid code", code: "NZ", want: "Oceania", wantErr: false},
		{name: "Valid code", code: "JP", want: "Asia", wantErr: false},
		{name: "Invalid code", code: "XX", want: "", wantErr: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CountryGetContinent(tc.code)
			if got != tc.want {
				t.Errorf("CountryGetContinent(%s) = %s; want %s", tc.code, got, tc.want)
			}
		})
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
