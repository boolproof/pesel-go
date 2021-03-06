package pesel

import (
	"fmt"
	"testing"
)

type testCase struct {
	title        string
	number       string
	err          bool
	expNumber    string
	expGender    string
	expBirthDate string
}

func TestNewPesel(t *testing.T) {
	testCases := []testCase{
		{"code too long", "700101001241", true, "", "", "0000-00-00"},
		{"code too long 2", "43221770011100121234", true, "", "", "0000-00-00"},
		{"code too short", "7001010012", true, "", "", "0000-00-00"},
		{"code too short 2", "700101", true, "", "", "0000-00-00"},
		{"invalid year", "7A010100123", true, "", "", "0000-00-00"},
		{"invalid month", "70A10100123", true, "", "", "0000-00-00"},
		{"invalid day", "70010.00123", true, "", "", "0000-00-00"},
		{"invalid date", "70000100123", true, "", "", "0000-00-00"},
		{"invalid date 2", "70130100129", true, "", "", "0000-00-00"},
		{"invalid date 3", "70010000127", true, "", "", "0000-00-00"},
		{"invalid date 3", "70013200128", true, "", "", "0000-00-00"},
		{"invalid date 4", "70022900129", true, "", "", "0000-00-00"},
		{"invalid checkSum 1", "70010100120", true, "", "", "0000-00-00"},
		{"invalid checkSum 2", "70010100121", true, "", "", "0000-00-00"},
		{"invalid checkSum 3", "70010100122", true, "", "", "0000-00-00"},
		{"invalid checkSum 4", "70010100123", true, "", "", "0000-00-00"},
		{"invalid checkSum 5", "70010100125", true, "", "", "0000-00-00"},
		{"invalid checkSum 6", "70010100126", true, "", "", "0000-00-00"},
		{"invalid checkSum 7", "70010100127", true, "", "", "0000-00-00"},
		{"invalid checkSum 8", "70010100128", true, "", "", "0000-00-00"},
		{"invalid checkSum 9", "70010100129", true, "", "", "0000-00-00"},
		{"correct 1", "70010100124", false, "70010100124", "female", "1970-01-01"},
		{"correct 2", "70010100193", false, "70010100193", "male", "1970-01-01"},
		{"correct 3", "70010101095", false, "70010101095", "male", "1970-01-01"},
		{"correct 4", "70010110097", false, "70010110097", "male", "1970-01-01"},
		{"correct 5", "70010199991", false, "70010199991", "male", "1970-01-01"},
		{"correct 6", "70123100125", false, "70123100125", "female", "1970-12-31"},
		{"correct 7", "69123100105", false, "69123100105", "female", "1969-12-31"},
		{"correct 2000s", "70210100120", false, "70210100120", "female", "2070-01-01"},
		{"correct 1800s", "01923100123", false, "01923100123", "female", "1801-12-31"},
		{"correct 2100s", "00410100116", false, "00410100116", "male", "2100-01-01"},
		{"correct 2200s", "99723100179", false, "99723100179", "male", "2299-12-31"},
		{"correct wiki 1", "55030101193", false, "55030101193", "male", "1955-03-01"},
		{"correct wiki 2", "55030101230", false, "55030101230", "male", "1955-03-01"},
	}

	for _, tc := range testCases {
		p, e := NewPesel(tc.number)
		if tc.err {
			if e == nil {
				t.Errorf("error should not be nil (%s)", tc.title)
			}
		} else {
			if e != nil {
				t.Errorf("error should be nil (%s)", tc.title)
			}
		}
		bd := fmt.Sprintf("%04d-%02d-%02d", p.BirthDate().Year, p.BirthDate().Month, p.BirthDate().Day)
		if bd != tc.expBirthDate {
			t.Errorf("expected birth date: %s, got: %s", tc.expBirthDate, bd)
		}
		if p.Number() != tc.expNumber {
			t.Errorf("expected code: %s, got: %s", tc.expNumber, p.Number())
		}
		if p.Gender() != tc.expGender {
			t.Errorf("expected gender: %s, got: %s", tc.expGender, p.Gender())
		}
	}
}
