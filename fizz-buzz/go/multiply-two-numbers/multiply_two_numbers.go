package multiplytwonumbers

import "fmt"

func main() {
	fmt.Println("Multiply two integrer numbers without the * operator")

}

// 2 * 3 = 6 i.e 2 + 2 + 2
// 2 * -3 = -6 i.e -2 + -2 + -2

func Multiply(x, y int) int {
	if x < 0 {
		x = -x
		y = -y
	}

	if y < 0 {
		x, y = y, x
	}

	sum := 0
	for i := 0; i < y; i++ {
		sum += x
	}
	return sum
}
