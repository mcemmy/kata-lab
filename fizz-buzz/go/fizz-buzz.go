package main

import "fmt"

func main() {
	var numbers = FizzBuzz()
	fmt.Println("Hello, Fizz-buzz!")
	fmt.Println("********************")
	fmt.Println(numbers)
}

//Write a program that prints the numbers from 1 to 100.
//But for multiples of three print "Fizz" instead of the number and for the multiples of five print "Buzz".
//For numbers which are multiples of both three and five print "FizzBuzz".

func FizzBuzz() []string {

	result := make([]string, 100)
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			result[i-1] = "FizzBuzz"
			continue
		}
		if i%3 == 0 {
			result[i-1] = "Fizz"
			continue
		} else if i%5 == 0 {
			result[i-1] = "Buzz"
			continue
		} else {
			result[i-1] = fmt.Sprint(i)
		}

	}
	return result
}
