package main

import "fmt"

func main() {
	// all primitives have a zero type (instantiated but not set)


	// ----------------- Booleans ------------------ //
	var a bool = true
	fmt.Printf("%v, %T\n", a, a)
	// will show up as a zero type (false)
	var b bool
	fmt.Printf("%v, %T\n", b, b)
	
	// ----------------- Integers ------------------ //
	fmt.Println("Integers")
	// zero type
	var zeroType int8
	fmt.Printf("%v, %T\n", zeroType, zeroType)
	
	// int range should be 64 or 128 (I think)
	c := 42
	fmt.Printf("%v, %T\n", c, c)
	
	// int8 range (-128 | 127)
	var smallInt int8
	fmt.Printf("%v, %T\n", smallInt, smallInt)
	
	// int16 range (-32768 | 32767)
	var medInt int16
	fmt.Printf("%v, %T\n", medInt, medInt)
	
	// int32 range (-2,147,483,648 | 2,147,483,647)
	var largeInt int32
	fmt.Printf("%v, %T\n", largeInt, largeInt)
	
	// int64 range (-9,223,372,036,854,775,808 | 9,223,372,036,854,775,807)
	var extraLargeInt int64
	fmt.Printf("%v, %T\n", extraLargeInt, extraLargeInt)

	// ! Important, cannot divide an integer and get the remainder, you must use % (modulo) to get the remainder if dividing two integers

		// ----------------- Floats ------------------ //
		fmt.Println("Floats")

		// zero type
		var zeroTypeFloat float64 // zeroTypeFloat == 0
		fmt.Printf("%v, %T\n", zeroTypeFloat, zeroTypeFloat)

		// float
		float1 := 3.14
		// float64 range (+- 2.23e-308 to +- 1.8E308)
		float1 := 3.14

	
	// ----------------- Strings ------------------ //
	fmt.Println("Strings")
	
	

	
	// ----------------- bit operators ------------------ //
	fmt.Println("Bitwise operators")
	d := 10
	e := 3
	fmt.Println(d & e) // and: d & e => (1010 & 0011) => 0010 (2 in binary)
	fmt.Println(d | e) // or: d & e => (1010 & 0011) => 1011 (11 in binary)
	fmt.Println(d ^ e) // xor: d & e => (1010 & 0011) => 1001 (9 in binary)
	fmt.Println(d &^ e) // xnot: d & e => (1010 & 0011) => 1000 (8 in binary)
	f := 8
	fmt.Println(f << 1) // Bit Shifting: shift f left 1 place == 16
	fmt.Println(f >> 1) // Bit Shifting: shift f right 1 place == 4

}