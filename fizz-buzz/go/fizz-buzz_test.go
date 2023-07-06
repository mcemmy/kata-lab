package main

import "testing"

//Write a program that prints the numbers from 1 to 100.
//But for multiples of three print "Fizz" instead of the number
//and for the multiples of five print "Buzz".
//For numbers which are multiples of both three and five print "FizzBuzz".
func TestPrintsNumberFromOneToHundred(t *testing.T) {
	result := FizzBuzz()
	//assert that the result is 100

	if len(result) != 100 {
		t.Errorf("Expected 100 results, got %d", len(result))
	}

	if result[0] != "1" {
		t.Errorf("Expected 1, got %s", result[0])
	}
}

func TestPrintsFizzForMultiplesOfThree(t *testing.T) {
	result := FizzBuzz()
	if result[2] != "Fizz" {
		t.Errorf("Expected Fizz, got %s", result[2])
	}
	if result[5] != "Fizz" {
		t.Errorf("Expected Fizz, got %s", result[5])
	}
	if result[98] != "Fizz" {
		t.Errorf("Expected Fizz, got %s", result[98])
	}
}

func TestPrintsBuzzForMultiplesOfFive(t *testing.T) {
	result := FizzBuzz()
	if result[4] != "Buzz" {
		t.Errorf("Expected Buzz, got %s", result[4])
	}
	if result[9] != "Buzz" {
		t.Errorf("Expected Buzz, got %s", result[9])
	}
	if result[19] != "Buzz" {
		t.Errorf("Expected Buzz, got %s", result[19])
	}
	if result[99] != "Buzz" {
		t.Errorf("Expected Buzz, got %s", result[99])
	}
}

func TestPrintsFizzBuzzForMultiplesOfThreeAndFive(t *testing.T) {
	result := FizzBuzz()
	if result[14] != "FizzBuzz" {
		t.Errorf("Expected FizzBuzz, got %s", result[14])
	}
	if result[29] != "FizzBuzz" {
		t.Errorf("Expected FizzBuzz, got %s", result[29])
	}
	if result[44] != "FizzBuzz" {
		t.Errorf("Expected FizzBuzz, got %s", result[44])
	}
	if result[59] != "FizzBuzz" {
		t.Errorf("Expected FizzBuzz, got %s", result[59])
	}
	if result[74] != "FizzBuzz" {
		t.Errorf("Expected FizzBuzz, got %s", result[74])
	}
	if result[89] != "FizzBuzz" {
		t.Errorf("Expected FizzBuzz, got %s", result[89])
	}
}
