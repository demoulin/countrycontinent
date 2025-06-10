# Country Continent

[![Go Report Card](https://goreportcard.com/badge/github.com/demoulin/countrycontinent)](https://goreportcard.com/report/github.com/demoulin/countrycontinent)
[![codecov](https://codecov.io/gh/demoulin/countrycontinent/graph/badge.svg?token=MEFFJJBA82)](https://codecov.io/gh/demoulin/countrycontinent)
[![GoDoc](https://godoc.org/github.com/demoulin/countrycontinent?status.svg)](https://godoc.org/github.com/demoulin/countrycontinent)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/demoulin/countrycontinent/blob/master/LICENSE)

The `country-continent` package provides a convenient way to retrieve information about countries and continents. It allows you to map country codes to country names and continents, as well as retrieve countries by continent.

## Features

- Get the full name of a country from its country code
- Get the full name and continent of a country from its country code
- Get the continent of a country from its country code
- Get a list of country codes belonging to a specific continent

## Installation

To install the package, use the following command:

```shell
go get github.com/demoulin/countrycontinent@v1.5.1
```

## Usage

### Get full name of a country

```go
func CountryGetFullName(countryCode string) string
```

Returns the full name of a country given its country code.

### Get full name and continent of a country

```go
func CountryGetFullNameContinent(countryCode string) (string, string)
```

Returns the full name and continent of a country given its country code.

### Get continent of a country

```go
func CountryGetContinent(countryCode string) string
```

Returns the continent of a country given its country code.

### Get countries in a continent

```go
func ContinentGetCountries(continent string) []string
```

Returns a list of country codes belonging to a given continent.

## Example

```go
package main

import (
    "fmt"
    "github.com/demoulin/countrycontinent"
)

func main() {
    // Get the full name of a country
    fmt.Println(countrycontinent.CountryGetFullName("US"))  // Output: "United States"

    // Get the full name and continent of a country
    name, continent := countrycontinent.CountryGetFullNameContinent("US")
    fmt.Printf("Country: %s, Continent: %s\n", name, continent)  // Output: Country: United States, Continent: North America

    // Get the continent of a country
    fmt.Println(countrycontinent.CountryGetContinent("US"))  // Output: "North America"

    // Get countries in a continent
    countries := countrycontinent.ContinentGetCountries("Europe")
    fmt.Println(countries)  // Output: [AD AL AT BA BE BG BY CH CZ DE DK EE ES FI FO FR GB GI GR HR HU IE IS IT LI LT LU LV MC MD MK MT NL NO PL PT RO RU SE SI SJ SK SM UA VA YU]
}
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This package is licensed under the [MIT License](https://github.com/demoulin/countrycontinent/blob/master/LICENSE).

---

Forked from <https://github.com/ArsFy/countrycontinent>
