package countrycontinent

import (
	"context"
)

// CountryInfo contains information about a country
type CountryInfo struct {
	Code      string
	Name      string
	Continent string
}

// QueryResult represents a generic result from a query operation
type QueryResult[T any] struct {
	Data  T
	Error error
}

// CountryQuery is a generic query function for countries
type CountryQuery[T any] func(ctx context.Context, countryCode string, opts *QueryOptions) QueryResult[T]

// CountryNameQuery returns a query function that retrieves just the country name
func CountryNameQuery() CountryQuery[string] {
	return func(ctx context.Context, countryCode string, opts *QueryOptions) QueryResult[string] {
		name, err := CountryGetFullNameWithContext(ctx, countryCode, opts)
		return QueryResult[string]{
			Data:  name,
			Error: err,
		}
	}
}

// CountryContinentQuery returns a query function that retrieves just the continent
func CountryContinentQuery() CountryQuery[string] {
	return func(ctx context.Context, countryCode string, opts *QueryOptions) QueryResult[string] {
		continent, err := CountryGetContinentWithContext(ctx, countryCode, opts)
		return QueryResult[string]{
			Data:  continent,
			Error: err,
		}
	}
}

// CountryFullInfoQuery returns a query function that retrieves all country information
func CountryFullInfoQuery() CountryQuery[CountryInfo] {
	return func(ctx context.Context, countryCode string, opts *QueryOptions) QueryResult[CountryInfo] {
		name, continent, err := CountryGetFullNameContinentWithContext(ctx, countryCode, opts)
		return QueryResult[CountryInfo]{
			Data: CountryInfo{
				Code:      countryCode,
				Name:      name,
				Continent: continent,
			},
			Error: err,
		}
	}
}

// ContinentCountriesQuery returns a query function that retrieves all countries in a continent
func ContinentCountriesQuery() func(ctx context.Context, continent string, opts *QueryOptions) QueryResult[[]string] {
	return func(ctx context.Context, continent string, opts *QueryOptions) QueryResult[[]string] {
		countries, err := ContinentGetCountriesWithContext(ctx, continent, opts)
		return QueryResult[[]string]{
			Data:  countries,
			Error: err,
		}
	}
}

// MultiCountryQuery executes a query against multiple country codes
func MultiCountryQuery[T any](ctx context.Context, countryCodes []string, query CountryQuery[T], opts *QueryOptions) []QueryResult[T] {
	results := make([]QueryResult[T], len(countryCodes))
	
	for i, code := range countryCodes {
		results[i] = query(ctx, code, opts)
	}
	
	return results
}

// FilterCountries filters countries by a predicate function
func FilterCountries[T any](results []QueryResult[T], predicate func(QueryResult[T]) bool) []QueryResult[T] {
	filtered := make([]QueryResult[T], 0)
	
	for _, result := range results {
		if predicate(result) {
			filtered = append(filtered, result)
		}
	}
	
	return filtered
}

// MapResults transforms query results using a mapping function
func MapResults[T any, R any](results []QueryResult[T], mapper func(QueryResult[T]) R) []R {
	mapped := make([]R, len(results))
	
	for i, result := range results {
		mapped[i] = mapper(result)
	}
	
	return mapped
}