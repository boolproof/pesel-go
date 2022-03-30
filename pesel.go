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

func NewPesel(code string) (Pesel, error) {
	p := Pesel{}
	e := errors.New("invalid PESEL")

	if len(code) != 11 {
		return p, e
	}

	ws := []int{1, 3, 7, 9, 1, 3, 7, 9, 1, 3, 1}
	ds := make([]int, 0, 11)
	var sum int

	for i := 0; i < 11; i++ {
		ch := string(code[i])
		d, err := strconv.Atoi(ch)
		if err != nil {
			return p, e
		}
		ds = append(ds, d)
		sum += d * ws[i]
	}

	if sum%10 != 0 {
		return p, e
	}

	y, m := 10*ds[0]+ds[1], 10*ds[2]+ds[3]

	var year int
	mod := m / 20
	month := m - mod*20
	switch mod {
	case 0:
		year = 1900 + y
	case 1:
		year = 2000 + y
	case 2:
		year = 2100 + y
	case 3:
		year = 2200 + y
	case 4:
		year = 1800 + y
	}

	bd, err := time.Parse("20060102", fmt.Sprintf("%04d%02d%02d", year, month, 10*ds[4]+ds[5]))

	if err != nil || 100*ds[6]+10*ds[7]+ds[8] == 0 {
		return p, e
	}

	p.code = code
	p.birthDate = &bd

	if ds[9]%2 == 0 {
		p.gender = GenderFemale
	} else {
		p.gender = GenderMale
	}

	return p, nil
}
