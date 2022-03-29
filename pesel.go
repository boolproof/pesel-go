package pesel

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const (
	GenderFemale     = "female"
	GenderMale       = "male"
	InvalidCodeError = "invalid PESEL"
)

type Pesel struct {
	code      string
	gender    string
	birthDate *time.Time
}

func (p Pesel) Code() string {
	return p.code
}

func (p Pesel) Gender() string {
	return p.gender
}

func (p Pesel) BirthDate() *time.Time {
	return p.birthDate
}

type parts struct {
	year     string
	month    string
	day      string
	sequence string
	gender   string
	checkSum string
}

type intParts struct {
	year     int
	month    int
	day      int
	sequence int
	gender   int
	checkSum int
}

func NewPesel(code string) (Pesel, error) {
	pesel := Pesel{}
	e := errors.New(InvalidCodeError)

	parts, err := newParts(code)
	if err != nil {
		return pesel, e
	}

	intParts, err := newIntParts(parts)
	if err != nil {
		return pesel, e
	}

	var year, month, centuryMod int
	centuryMod = intParts.month / 20
	month = intParts.month - centuryMod*20
	switch centuryMod {
	case 0:
		year = 1900 + intParts.year
	case 1:
		year = 2000 + intParts.year
	case 2:
		year = 2100 + intParts.year
	case 3:
		year = 2200 + intParts.year
	case 4:
		year = 1800 + intParts.year
	}

	layout := "2006-01-02"
	str := fmt.Sprintf("%04d-%02d-%02d", year, month, intParts.day)
	birthDate, err := time.Parse(layout, str)

	if err != nil {
		return pesel, e
	}

	if intParts.sequence == 0 {
		return pesel, e
	}

	if !validateCheckSum(code) {
		return pesel, e
	}

	pesel.code = code
	pesel.birthDate = &birthDate

	if intParts.gender%2 == 0 {
		pesel.gender = GenderFemale
	} else {
		pesel.gender = GenderMale
	}

	return pesel, nil
}

func validateCheckSum(code string) bool {
	w := []int{1, 3, 7, 9, 1, 3, 7, 9, 1, 3, 1}
	sum := 0
	for i := 0; i < 11; i++ {
		d, _ := strconv.Atoi(code[i : i+1])
		sum += w[i] * d
	}

	if sum%10 == 0 {
		return true
	}

	return false
}

func newParts(code string) (parts, error) {
	if len(code) != 11 {
		return parts{}, errors.New(InvalidCodeError)
	}

	return parts{
		year:     code[0:2],
		month:    code[2:4],
		day:      code[4:6],
		sequence: code[6:9],
		gender:   code[9:10],
		checkSum: code[10:11],
	}, nil
}

func newIntParts(parts parts) (intParts, error) {
	year, err := strconv.Atoi(parts.year)
	if err != nil {
		return intParts{}, err
	}

	month, err := strconv.Atoi(parts.month)
	if err != nil {
		return intParts{}, err
	}

	day, err := strconv.Atoi(parts.day)
	if err != nil {
		return intParts{}, err
	}

	sequence, err := strconv.Atoi(parts.sequence)
	if err != nil {
		return intParts{}, err
	}

	gender, err := strconv.Atoi(parts.gender)
	if err != nil {
		return intParts{}, err
	}

	checkSum, err := strconv.Atoi(parts.checkSum)
	if err != nil {
		return intParts{}, err
	}

	intParts := intParts{
		year:     year,
		month:    month,
		day:      day,
		sequence: sequence,
		gender:   gender,
		checkSum: checkSum,
	}

	return intParts, nil
}
