package pesel

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const (
	GenderFemale = "female"
	GenderMale   = "male"
)

type Pesel struct {
	code      string
	gender    string
	birthDate *Date
}

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

func (p Pesel) Code() string {
	return p.code
}

func (p Pesel) Gender() string {
	return p.gender
}

func (p Pesel) BirthDate() *Date {
	return p.birthDate
}

func NewPesel(code string) (Pesel, error) {
	p := Pesel{}
	e := errors.New("invalid PESEL")

	if len(code) != 11 {
		return p, e
	}

	ws := [11]int{1, 3, 7, 9, 1, 3, 7, 9, 1, 3, 1}
	ds := [11]int{}
	var sum int

	for i := 0; i < 11; i++ {
		d, err := strconv.Atoi(string(code[i]))
		if err != nil {
			return p, e
		}
		ds[i] = d
		sum += d * ws[i]
	}

	if sum%10 != 0 {
		return p, e
	}

	m := 10*ds[2] + ds[3]
	mod := m / 20
	cs := [5]int{1900, 2000, 2100, 2200, 1800}
	bd, err := time.Parse("20060102", fmt.Sprintf("%04d%02d%02d", cs[mod]+10*ds[0]+ds[1], m-mod*20, 10*ds[4]+ds[5]))

	if err != nil || 100*ds[6]+10*ds[7]+ds[8] == 0 {
		return p, e
	}

	p.code = code
	p.birthDate = &Date{
		Year:  bd.Year(),
		Month: bd.Month(),
		Day:   bd.Day(),
	}

	if ds[9]%2 == 0 {
		p.gender = GenderFemale
	} else {
		p.gender = GenderMale
	}

	return p, nil
}
