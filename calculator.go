package main

import "fmt"

func addition(a, b int) (result int) {
	result = a + b
	return result
}

func subtract(a, b int) (result int) {
	result = a - b
	return result
}

func multiply(a, b int) (result int) {
	result = a * b
	return result
}

func divide(a, b int) (result int) {
	result = a / b
	return result
}

func main() {
	result := addition(5, 5)
	fmt.Println(result)

	result = subtract(5, 3)
	fmt.Println(result)

	result = multiply(5, 8)
	fmt.Println(result)

	result = divide(20, 5)
	fmt.Println(result)
}
