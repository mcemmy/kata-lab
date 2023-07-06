package main

import "testing"

//Write a program that prints the numbers from 1 to 100.
//But for multiples of three print "Fizz" instead of the number and for the multiples of five print "Buzz".
//For numbers which are multiples of both three and five print "FizzBuzz".
func TestPrintsNumberFromOneToHundred(t *testing.T) {
	result := FizzBuzz()
	//assert that the result is 100

	if len(result) != 100 {
		t.Errorf("Expected 100 results, got %d", len(result))
	}

	if result[0] != 1 {
		t.Errorf("Expected 1, got %d", result[0])
	}

	if result[99] != 100 {
		t.Errorf("Expected 100, got %d", result[99])
	}
}
