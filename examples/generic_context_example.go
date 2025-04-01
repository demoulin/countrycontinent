package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/demoulin/countrycontinent/v1.5.1"
)

func main() {
	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create query options with timeout and tags
	opts := countrycontinent.DefaultQueryOptions().
		WithTimeout(2 * time.Second).
		WithTag("request_id", "example-123").
		WithTag("source", "example-app")

	// Using context-aware functions directly
	name, err := countrycontinent.CountryGetFullNameWithContext(ctx, "US", opts)
	if err != nil {
		log.Fatalf("Error getting country name: %v", err)
	}
	fmt.Printf("Country name: %s\n", name)

	// Create a generic query for country info
	infoQuery := countrycontinent.CountryFullInfoQuery()

	// Query a single country
	result := infoQuery(ctx, "FR", opts)
	if result.Error != nil {
		log.Fatalf("Error in query: %v", result.Error)
	}
	fmt.Printf("Country info: %+v\n", result.Data)

	// Batch query multiple countries
	countryCodes := []string{"DE", "IT", "JP", "BR", "XX"}
	results := countrycontinent.MultiCountryQuery(ctx, countryCodes, infoQuery, opts)

	// Filter for successful results only
	successfulResults := countrycontinent.FilterCountries(results, func(r countrycontinent.QueryResult[countrycontinent.CountryInfo]) bool {
		return r.Error == nil
	})

	// Map to just country names
	countryNames := countrycontinent.MapResults(successfulResults, func(r countrycontinent.QueryResult[countrycontinent.CountryInfo]) string {
		return r.Data.Name
	})

	fmt.Println("Valid country names from batch:")
	for _, name := range countryNames {
		fmt.Printf("- %s\n", name)
	}

	// Find all countries in Asia
	continentQuery := countrycontinent.ContinentCountriesQuery()
	asiaResult := continentQuery(ctx, "Asia", opts)
	if asiaResult.Error != nil {
		log.Fatalf("Error getting Asian countries: %v", asiaResult.Error)
	}

	fmt.Printf("Found %d countries in Asia\n", len(asiaResult.Data))

	// Demonstrate timeout handling
	timeoutCtx, timeoutCancel := context.WithTimeout(ctx, 1*time.Nanosecond)
	defer timeoutCancel()

	// This should fail with timeout
	_, timeoutErr := countrycontinent.CountryGetFullNameWithContext(timeoutCtx, "US", opts)
	fmt.Printf("Expected timeout error: %v\n", timeoutErr)
}