package countrycontinent

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

// TestCountryGetFullName tests the CountryGetFullName function.
func TestCountryGetFullName(t *testing.T) {
	tests := []struct {
		name          string
		code          string
		expected      string
		expectedError error
	}{
		{name: "Valid code US", code: "US", expected: "United States", expectedError: nil},
		{name: "Unknown code XX", code: "XX", expected: "", expectedError: &CountryNotFoundError{CountryCode: "XX"}},
		{name: "Empty code", code: "", expected: "", expectedError: &InvalidCountryCodeError{CountryCode: ""}},
		{name: "Lowercase code fr", code: "fr", expected: "", expectedError: &InvalidCountryCodeError{CountryCode: "fr"}},
		{name: "Lowercase code xx", code: "xx", expected: "", expectedError: &InvalidCountryCodeError{CountryCode: "xx"}},
		{name: "Mixed case code Xx", code: "Xx", expected: "", expectedError: &InvalidCountryCodeError{CountryCode: "Xx"}},
		{name: "Valid code AD", code: "AD", expected: "Andorra", expectedError: nil},
		{name: "Valid code ZA", code: "ZA", expected: "South Africa", expectedError: nil},
		{name: "Valid code ZW", code: "ZW", expected: "Zimbabwe", expectedError: nil},
		{name: "Invalid format - too short", code: "A", expected: "", expectedError: &InvalidCountryCodeError{CountryCode: "A"}},
		{name: "Invalid format - too long", code: "USA", expected: "", expectedError: &InvalidCountryCodeError{CountryCode: "USA"}},
		{name: "Invalid format - contains number", code: "U1", expected: "", expectedError: &InvalidCountryCodeError{CountryCode: "U1"}},
		{name: "Invalid format - contains special char", code: "U@", expected: "", expectedError: &InvalidCountryCodeError{CountryCode: "U@"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CountryGetFullName(tc.code)
			if !reflect.DeepEqual(err, tc.expectedError) {
				t.Errorf("CountryGetFullName(%s) error = %v, wantError %v", tc.code, err, tc.expectedError)
			}
			if err == nil && got != tc.expected { // Only check 'got' if no error was expected
				t.Errorf("CountryGetFullName(%s) got = %v, want %v", tc.code, got, tc.expected)
			}
		})
	}
}

func TestErrorMessages(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name:     "InvalidCountryCodeError",
			err:      &InvalidCountryCodeError{CountryCode: "XX"},
			expected: "invalid country code format: XX",
		},
		{
			name:     "CountryNotFoundError",
			err:      &CountryNotFoundError{CountryCode: "YY"},
			expected: "country code not found: YY",
		},
		{
			name:     "ContinentNotFoundError",
			err:      &ContinentNotFoundError{Continent: "Mars"},
			expected: "continent not found: Mars",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.err.Error() != tc.expected {
				t.Errorf("Error() got %q, want %q", tc.err.Error(), tc.expected)
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
		expectedError error
	}{
		{name: "Valid code NL", code: "NL", wantCountry: "Netherlands", wantContinent: "Europe", expectedError: nil},
		{name: "Valid code US", code: "US", wantCountry: "United States", wantContinent: "North America", expectedError: nil},
		{name: "Valid code MX", code: "MX", wantCountry: "Mexico", wantContinent: "North America", expectedError: nil},
		{name: "Valid code PE", code: "PE", wantCountry: "Peru", wantContinent: "South America", expectedError: nil},
		{name: "Valid code DM", code: "DM", wantCountry: "Dominica", wantContinent: "Caribbean", expectedError: nil},
		{name: "Valid code NZ", code: "NZ", wantCountry: "New Zealand (Aotearoa)", wantContinent: "Oceania", expectedError: nil},
		{name: "Valid code JP", code: "JP", wantCountry: "Japan", wantContinent: "Asia", expectedError: nil},
		{name: "Unknown code XX", code: "XX", wantCountry: "", wantContinent: "", expectedError: &CountryNotFoundError{CountryCode: "XX"}},
		{name: "Valid code AD", code: "AD", wantCountry: "Andorra", wantContinent: "Europe", expectedError: nil},
		{name: "Valid code ZA", code: "ZA", wantCountry: "South Africa", wantContinent: "Africa", expectedError: nil},
		{name: "Valid code ZW", code: "ZW", wantCountry: "Zimbabwe", wantContinent: "Africa", expectedError: nil},
		{name: "Empty code", code: "", wantCountry: "", wantContinent: "", expectedError: &InvalidCountryCodeError{CountryCode: ""}},
		{name: "Lowercase code fr", code: "fr", wantCountry: "", wantContinent: "", expectedError: &InvalidCountryCodeError{CountryCode: "fr"}},
		{name: "Invalid format - too short", code: "A", wantCountry: "", wantContinent: "", expectedError: &InvalidCountryCodeError{CountryCode: "A"}},
		{name: "Invalid format - too long", code: "USA", wantCountry: "", wantContinent: "", expectedError: &InvalidCountryCodeError{CountryCode: "USA"}},
		{name: "Invalid format - contains number", code: "U1", wantCountry: "", wantContinent: "", expectedError: &InvalidCountryCodeError{CountryCode: "U1"}},
		{name: "Invalid format - contains special char", code: "U@", wantCountry: "", wantContinent: "", expectedError: &InvalidCountryCodeError{CountryCode: "U@"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotCountry, gotContinent, err := CountryGetFullNameContinent(tc.code)
			if !reflect.DeepEqual(err, tc.expectedError) {
				t.Errorf("CountryGetFullNameContinent(%s) error = %v, wantError %v", tc.code, err, tc.expectedError)
			}
			if err == nil && (gotCountry != tc.wantCountry || gotContinent != tc.wantContinent) {
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
		name          string
		code          string
		want          string
		expectedError error
	}{
		{name: "Valid code NL", code: "NL", want: "Europe", expectedError: nil},
		{name: "Valid code US", code: "US", want: "North America", expectedError: nil},
		{name: "Valid code MX", code: "MX", want: "North America", expectedError: nil},
		{name: "Valid code PE", code: "PE", want: "South America", expectedError: nil},
		{name: "Valid code DM", code: "DM", want: "Caribbean", expectedError: nil},
		{name: "Valid code NZ", code: "NZ", want: "Oceania", expectedError: nil},
		{name: "Valid code JP", code: "JP", want: "Asia", expectedError: nil},
		{name: "Unknown code XX", code: "XX", want: "", expectedError: &CountryNotFoundError{CountryCode: "XX"}},
		{name: "Valid code AD", code: "AD", want: "Europe", expectedError: nil},
		{name: "Valid code ZA", code: "ZA", want: "Africa", expectedError: nil},
		{name: "Valid code ZW", code: "ZW", want: "Africa", expectedError: nil},
		{name: "Empty code", code: "", want: "", expectedError: &InvalidCountryCodeError{CountryCode: ""}},
		{name: "Lowercase code fr", code: "fr", want: "", expectedError: &InvalidCountryCodeError{CountryCode: "fr"}},
		{name: "Invalid format - too short", code: "A", want: "", expectedError: &InvalidCountryCodeError{CountryCode: "A"}},
		{name: "Invalid format - too long", code: "USA", want: "", expectedError: &InvalidCountryCodeError{CountryCode: "USA"}},
		{name: "Invalid format - contains number", code: "U1", want: "", expectedError: &InvalidCountryCodeError{CountryCode: "U1"}},
		{name: "Invalid format - contains special char", code: "U@", want: "", expectedError: &InvalidCountryCodeError{CountryCode: "U@"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CountryGetContinent(tc.code)
			if !reflect.DeepEqual(err, tc.expectedError) {
				t.Errorf("CountryGetContinent(%s) error = %v, wantError %v", tc.code, err, tc.expectedError)
			}
			if err == nil && got != tc.want {
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

func TestIsValidCountryCode(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		expected bool
	}{
		{name: "Valid code - US", code: "US", expected: true},
		{name: "Valid code - DE", code: "DE", expected: true},
		{name: "Invalid code - empty", code: "", expected: false},
		{name: "Invalid code - too short", code: "A", expected: false},
		{name: "Invalid code - too long", code: "USA", expected: false},
		{name: "Invalid code - lowercase", code: "us", expected: false},
		{name: "Invalid code - mixed case", code: "Us", expected: false},
		{name: "Invalid code - contains number", code: "U1", expected: false},
		{name: "Invalid code - contains special character", code: "U@", expected: false},
		{name: "Invalid code - number only", code: "12", expected: false},
		{name: "Invalid code - special chars only", code: "@#", expected: false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isValidCountryCode(tc.code)
			if got != tc.expected {
				t.Errorf("isValidCountryCode(%q) = %v, want %v", tc.code, got, tc.expected)
			}
		})
	}
}
