package main

import "fmt"

func main(){
	// Print statement
	fmt.Println("Hello Go!")

	// ============== Variables ============== //
	// Naming Conventions
		// 1. length of variable name should reflect how long variable will be used. (i.e. single letter should be for loops, words or multiple words for package level or greater) 
	// lowercase variables = scoped to the package
	// Uppercase(first letter) variables = globally scoped
	// Three levels of visibility for variabels: exported from package(globally available), package level and scoped level (inside of func)

	// ----- declaration ----- //
	// variable declaration that can be used in other scopes and changed 
	var i int
	i = 42
	fmt.Println(i)
	// single line declaration - useful if you need to hand type a variable if go won't have enough info to do it automatically (example: number starts as an int but should be a float)
	var a int = 42
	fmt.Println(a)
	// slick single line declaration
	b := 42
	fmt.Println(b)


	// Multiple variable declaration
	// var (
	// 	actorName string = "Elisabeth Sladen"
	// 	companion string = "Sarah Jane Smith"
	// 	doctorNumber int = 3
	// 	season int = 11
	// )
	// var (
	// 	counter int = 0
	// )

	// Variables must be used

	// ----- redeclaration and shadowing ----- //
	// can declare a variable outside of a func and use it within and can also have a variable of the same name within the scope of the variable.
	
	// visibility
	
	
	// naming conventions
	
	
	// type conversions
	

	// ==============  ============== //

}