# country-continent

Get continent name from Country code, Country name and Get Country code from continent name.

-----

## Install

``` shell
go get github.com/demoulin/countrycontinent@v1.3.0
```

## Use

### Get countries from continent name

``` go
func ContinentGetCountry(ct string) []string
```

### Get continent name from Country code

``` go
func CountryGetContinent(cc string) string
```

### Get Country full name from Country code

``` go
func CountryGetFullName(cc string) string
```

## Example

```go
package main

import (
    "github.com/demoulin/countrycontinent"
)

func main(){
    countrycontinent.CountryGetContinent("HK")   // "Asia"
    countrycontinent.CountryGetFullName("HK")    // "Hong Kong"
    countrycontinent.ContinentGetCountry("Asia")
    // [AF AM AZ BH BD BT BN KH CN CX CC CY GE HK IN ID IR IQ IL JP JO KZ KP KR KW KG LA LB MY MV MN MM NP OM PK PH QA RU SA SG LK SY TW TJ TH TR TM AE UZ VN YE]
}
```

Forked from <https://github.com/ArsFy/countrycontinent> 