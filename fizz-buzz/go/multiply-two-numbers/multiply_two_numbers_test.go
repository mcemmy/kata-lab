package multiplytwonumbers

import "testing"

func Test2MultiplyBy3ShouldBe6(t *testing.T) {

	result := Multiply(2, 3)

	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}

func Test3MultiplyBy3ShouldBe9(t *testing.T) {

	result := Multiply(3, 3)

	if result != 9 {
		t.Errorf("Expected 9, got %d", result)
	}
}

func TestNegative2MultiplyBy3ShouldBeNegative6(t *testing.T) {

	result := Multiply(-2, 3)

	if result != -6 {
		t.Errorf("Expected -6, got %d", result)
	}
}

func TestNegative2MultiplyByNegative3ShouldBe6(t *testing.T) {

	result := Multiply(-2, -3)

	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}

func Test2MultiplyByNegative3ShouldBeNegative6(t *testing.T) {

	result := Multiply(2, -3)

	if result != -6 {
		t.Errorf("Expected -6, got %d", result)
	}
}

func Test0MultiplyBy3ShouldBe0(t *testing.T) {

	result := Multiply(0, 3)

	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func Test0MultiplyByNegative3ShouldBe0(t *testing.T) {

	result := Multiply(0, -3)

	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
