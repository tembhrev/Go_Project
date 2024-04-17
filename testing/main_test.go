package main

import "testing"

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 2.0)
	if err != nil {
		t.Error("Got an error when should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0.0)
	if err == nil {
		t.Error("Dig not get an error when should have")
	}
}

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isError  bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, false},
}

func TestDivision(t *testing.T) {
	for _, test := range tests {
		got, err := divide(test.dividend, test.divisor)
		if test.isError {
			if err == nil {
				t.Error("Expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect an error but got one", err.Error())
			}
		}
		if got != test.expected {
			t.Errorf("Expected %f but got %f", test.expected, got)
		}
	}
}
