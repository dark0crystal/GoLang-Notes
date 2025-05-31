package main

import (
	"fmt"
)

// return one value
func add(x, y int) int {
	return x + y
}

// return multiple values
func add2(x, y int) (int, int) {
	return x + y, x * y
}

// labeled return values
func multi(x, y int) (num1 int, num2 int) {
	defer fmt.Println("Hello to the end of the world !!")
	num1 = x * y
	num2 = x*y + x
	fmt.Println("Before return...")
	return
}

func main() {
	sum := add(5, 4)
	fmt.Println("Return one value: ", sum)
	sum2, mult := add2(5, 4)
	fmt.Printf("The sum is %d , and the multiply is: %d \n", sum2, mult)
	num1, num2 := multi(4, 5)
	fmt.Printf("First Num: %d , The second Num : %d \n", num1, num2)
	// creating a function inside the main and pass it to a value is valid thing !!

	test := func() {
		fmt.Println("Hello form inside the main func :)")

	}
	// to call the function , you should add paranthesis to the end of the variable
	test()

}
