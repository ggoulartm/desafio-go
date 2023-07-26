package main

import "testing"

func TestShouldSumCorrect(t *testing.T) {
	test := sum(5, 10, 15)
	result := 30
	if test != result {
		t.Error("Expected value: ", result, " Returned value: ", test)
	}
}

func TestShouldSubCorrect(t *testing.T) {
	test := sub(15, 10)
	result := 5
	if test != result {
		t.Error("Expected value: ", result, " Returned value: ", test)
	}
}

func TestShouldMultCorrect(t *testing.T) {
	test := mult(15, 10, 2)
	result := 300
	if test != result {
		t.Error("Expected value: ", result, " Returned value: ", test)
	}
}

func TestShouldDivCorrect(t *testing.T) {
	test := divider(15, 3, 5)
	result := 1
	if test != result {
		t.Error("Expected value: ", result, " Returned value: ", test)
	}
}
