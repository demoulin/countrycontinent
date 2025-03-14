package countrycontinent

import (
	"sort"
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
		{name: "Andorra", code: "AD", expected: "Andorra", wantErr: false},
		{name: "South Africa", code: "ZA", expected: "South Africa", wantErr: false},
		{name: "Zimbabwe", code: "ZW", expected: "Zimbabwe", wantErr: false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CountryGetFullName(tc.code)
			if (err != nil) != tc.wantErr {
				t.Errorf("CountryGetFullName(%s) error = %v, wantErr %v", tc.code, err, tc.wantErr)
				return
			}
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
		{name: "Andorra", code: "AD", wantCountry: "Andorra", wantContinent: "Europe", wantErr: false},
		{name: "South Africa", code: "ZA", wantCountry: "South Africa", wantContinent: "Africa", wantErr: false},
		{name: "Zimbabwe", code: "ZW", wantCountry: "Zimbabwe", wantContinent: "Africa", wantErr: false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotCountry, gotContinent, err := CountryGetFullNameContinent(tc.code)
			if (err != nil) != tc.wantErr {
				t.Errorf("CountryGetFullNameContinent(%s) error = %v, wantErr %v", tc.code, err, tc.wantErr)
				return
			}
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
		wantErr   bool
	}{
		{name: "Europe", continent: "Europe", wantErr: false},
		{name: "North America", continent: "North America", wantErr: false},
		{name: "Central America", continent: "Central America", wantErr: false},
		{name: "South America", continent: "South America", wantErr: false},
		{name: "Caribbean", continent: "Caribbean", wantErr: false},
		{name: "Oceania", continent: "Oceania", wantErr: false},
		{name: "Asia", continent: "Asia", wantErr: false},
		{name: "Africa", continent: "Africa", wantErr: false},
		{name: "Antarctica", continent: "Antarctica", wantErr: false},
		{name: "Invalid continent", continent: "XX", wantErr: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ContinentGetCountries(tc.continent)
			if (err != nil) != tc.wantErr {
				t.Errorf("ContinentGetCountries(%s) error = %v, wantErr %v", tc.continent, err, tc.wantErr)
				return
			}
			if tc.wantErr == false {
				expectedCountries, err := getCountriesByContinent(tc.continent)
				if err != nil {
					t.Errorf("getCountriesByContinent(%s) returned an error = %v", tc.continent, err)
				}
				if !equalSlices(got, expectedCountries) {
					t.Errorf("ContinentGetCountries(%s) = %s; want %s", tc.continent, got, expectedCountries)
				}
			} else {
				if err == nil {
					t.Errorf("ContinentGetCountries(%s) = %s; want an error", tc.continent, got)
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
		{name: "Andorra", code: "AD", want: "Europe", wantErr: false},
		{name: "South Africa", code: "ZA", want: "Africa", wantErr: false},
		{name: "Zimbabwe", code: "ZW", want: "Africa", wantErr: false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CountryGetContinent(tc.code)
			if (err != nil) != tc.wantErr {
				t.Errorf("CountryGetContinent(%s) error = %v, wantErr %v", tc.code, err, tc.wantErr)
				return
			}
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

func getCountriesByContinent(continent string) ([]string, error) {
	var countries []string
	for _, c := range countryContinent {
		if c.Continent == continent {
			countries = append(countries, c.CountryCode)
		}
	}
	if len(countries) == 0 {
		return nil, &ContinentNotFoundError{Continent: continent}
	}
	return countries, nil
}

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Strings(a)
	sort.Strings(b)
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
