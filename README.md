# PESEL number validator
Package provides validation of PESEL (the national identification number used in Poland - https://en.wikipedia.org/wiki/PESEL) for **Go**. If passed number is valid, result struct provides access to birthdate and gender extracted from the number while result error is `nil`. If number is invalid, result struct contains empty values (`nil` for birthdate) and the result error is not `nil`.

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
