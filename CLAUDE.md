# CLAUDE.md - Guidelines for Country-Continent Package

## Build/Lint/Test Commands
- Build: `go build`
- Run all tests: `go test -v ./...` 
- Run single test: `go test -v -run TestName` (e.g. `go test -v -run TestCountryGetFullName`)
- Run with coverage: `go test -cover ./...`
- Check formatting: `go fmt ./...`
- Run linter: `golint ./...`

## Code Style Guidelines
- **Formatting**: Follow standard Go formatting with `go fmt`
- **Imports**: Standard Go import style (stdlib first, then external)
- **Types**: Use strong typing and error returns (not just empty strings)
- **Naming**: 
  - Functions: CamelCase (e.g., `CountryGetFullName`)
  - Variables: camelCase (e.g., `countryCode`)
  - Constants: UPPERCASE
- **Error Handling**: Use custom error types with descriptive messages
- **Comments**: Document all exported functions with proper GoDoc format
- **Testing**: Create thorough tests with table-driven test cases
- **Code Organization**: Group related functions together