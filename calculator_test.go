package main

import "testing"

func TestShouldSumCorrect(t *testing.T) {
	test := sum(3, 2, 10)

	result := 15

	if test != result {
		t.Error("Valor esperado:", result, "Valor retornado:", test)
	}
}

func TestShouldMultiplyCorrect(t *testing.T) {
	test := multiply(10, 2)

	result := 20

	if test != result {
		t.Error("Valor esperado:", result, "Valor retornado:", test)
	}
}

func TestShouldSubtractCorrect(t *testing.T) {
	test := subtract(20, 18)

	result := 2

	if test != result {
		t.Error("Valor esperado:", result, "Valor retornado:", test)
	}
}

func TestShouldDivideCorrect(t *testing.T) {
	test := divide(10, 5, 2)

	result := 1

	if test != result {
		t.Error("Valor esperado:", result, "Valor retornado:", test)
	}
}
