package main

import (
	"fmt"
	"strconv"
)

// when declaring variables in package scope (here), you cannot use := syntax here, only inside a function

// you can declare a variable of the same name in a lower scope, but not in the same scope.
var packageScopeVar = "package Scope"

// Naming
// lowercase variables are scoped to the package
// uppercase first letter allows the variable to be exported from the package
// block scope variables are scoped to the block abd not visible outside of the block.
// Small variable names are for very locally used variables. Verbose names are for variables used outside of the package.
// Acronyms should be all uppercase
// Use camelcase

func main() {
	// Three ways to declare a variable
	var a int
	a = 1
	var b int = 2
	c := 3
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

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
	
	// ----------------- Scoping ------------------ //
	
	fmt.Println(packageScopeVar)
	packageScopeVar := "now func scope and is over shadowed"
	fmt.Println(packageScopeVar)

	// ----------------- Types ------------------ //
	// **Types must be explicitly converted
	// ----- float to int ----- //
	var d float32 = 42.5
	fmt.Println(d)
	// convert to integer and will lose the .5
	e := int(d) 
	fmt.Println(e)

	// ----- int to string ----- //
	var f int = 42
	fmt.Printf("%v, %T\n", f, f)
	
	var g string
	g = strconv.Itoa(f)
	fmt.Printf("%v, %T\n", g, g)

}