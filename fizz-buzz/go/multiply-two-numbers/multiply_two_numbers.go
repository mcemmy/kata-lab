package multiplytwonumbers

import "fmt"

func main() {
	fmt.Println("Hello, Multiply Two Numbers!")
	fmt.Println("********************")
}

func Multiply(x, y int) int {
	if x < 0 {
		x = -x
		y = -y
	}

	if y < 0 {
		x, y = y, x
	}

	result := 0
	for i := 0; i < y; i++ {
		result += x
	}
	return result
}
