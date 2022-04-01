[![Go Reference](https://pkg.go.dev/badge/github.com/boolproof/pesel-go.svg)](https://pkg.go.dev/github.com/boolproof/pesel-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/boolproof/pesel-go)](https://goreportcard.com/report/github.com/boolproof/pesel-go)

# PESEL number validator
Package provides validation of PESEL (the national identification number used in Poland - https://en.wikipedia.org/wiki/PESEL) for **Go**. If argument PESEL number (passed in as a `string`) is valid, result struct provides access to birthdate and gender extracted from the number, while result error is `nil`. If number is invalid, non `nil` error is returned, while the result struct contains zero values.

## Disclaimer
This package is not providing information if given PESEL number exists. Package only validates if the number complies with basic rules described by legislation. Positive validation doesn't mean given number has been issued and assigned to any person. Negative validation means that the number theoretically could have not been issued, unless by mistake of authorities.

## Example usage

```go
package main

import (
	"fmt"
	"os"

	"github.com/boolproof/pesel-go"
)

func main() {
	number := os.Args[1]

	pesel, err := pesel.NewPesel(number)
	if err != nil {
		fmt.Println(err.Error()) //outputs "invalid PESEL"
	} else {
		birthDateStr := fmt.Sprintf("%04d-%02d-%02d", pesel.BirthDate().Year, pesel.BirthDate().Month, pesel.BirthDate().Day)
		fmt.Printf("'%s' is a valid PESEL. Encoded birthdate: %s; encoded gender: %s\n", pesel.Number(), birthDateStr, pesel.Gender())
	}
}
```
