
// Function closures

// Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

// For example, the adder function returns a closure. Each closure is bound to its own sum variable.

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}


// Exercise: Fibonacci closure

// Let's have some fun with functions.

// Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	int1 := 0;
	int2 := 1;
	return func() int {
		sum := int1 + int2;
		int1 = int2
		int2 = sum
		return sum
		
	}
	
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}