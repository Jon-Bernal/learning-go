package main

import "fmt"

func types() {
	// all primitives have a zero type (instantiated but not set)

	// ----------------- Booleans ------------------ //
	fmt.Println("// ----- Booleans ----- //")
	var a bool = true
	fmt.Printf("%v, %T\n", a, a)
	// will show up as a zero type (false)
	var b bool
	fmt.Printf("%v, %T\n", b, b)

	// ----------------- Integers ------------------ //
	fmt.Println("// ----- Integers ----- //")
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

	// ----------------- bit operators ------------------ //
	fmt.Println("// ----- Bitwise operators ----- //")
	fmt.Println("Bitwise operators")
	d := 10
	e := 3
	fmt.Println(d & e)  // and: d & e => (1010 & 0011) => 0010 (2 in binary)
	fmt.Println(d | e)  // or: d & e => (1010 & 0011) => 1011 (11 in binary)
	fmt.Println(d ^ e)  // xor: d & e => (1010 & 0011) => 1001 (9 in binary)
	fmt.Println(d &^ e) // xnot: d & e => (1010 & 0011) => 1000 (8 in binary)
	f := 8
	fmt.Println(f << 1) // Bit Shifting: shift f left 1 place == 16
	fmt.Println(f >> 1) // Bit Shifting: shift f right 1 place == 4

	// ----------------- Floats ------------------ //
	fmt.Println("// ----- Floats ----- //")

	// zero type
	var zeroTypeFloat float64 // zeroTypeFloat == 0
	fmt.Printf("%v, %T\n", zeroTypeFloat, zeroTypeFloat)

	// float initialized like below will be a float64
	float1 := 3.14
	float1 = 13.7e72
	float1 = 2.1e14
	fmt.Printf("%v, %T\n", float1, float1)
	// float64 range (-3.4e+38 to 3.4e+38.)
	var float2 float32 = 3.14
	fmt.Printf("%v, %T\n", float2, float2)
	// float64 range (-1.7e+308 to +1.7e+308)
	var float3 float64 = 3.14
	fmt.Printf("%v, %T\n", float3, float3)

	// Float operations
	floatOp1 := 10.2
	floatOp2 := 3.7
	fmt.Println(floatOp1 + floatOp2)
	fmt.Println(floatOp1 - floatOp2)
	fmt.Println(floatOp1 * floatOp2)
	fmt.Println(floatOp1 / floatOp2)
	// Notes:
	// always returns a float
	// no remainder operator available
	// No bitwise operators for floats

	// ----------------- Complex Numbers ------------------ //
	fmt.Println("// ----- Complex Numbers ----- //")

	var c64 complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", c64, c64)

	// Operations
	complex1 := 1 + 2i
	complex2 := 2 + 5.2i
	fmt.Println(complex1 + complex2)
	fmt.Println(complex1 - complex2)
	fmt.Println(complex1 * complex2)
	fmt.Println(complex1 / complex2)

	fmt.Printf("%v, %T\n", real(complex1), real(complex1)) // returns just the real number part of imaginary number
	fmt.Printf("%v, %T\n", imag(complex2), imag(complex2)) // returns just the imaginary number part

	var newComplex complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", newComplex, newComplex) // returns just the imaginary number part

	// ----------------- Strings ------------------ //
	// Notes
	// Strings are immutable

	fmt.Println("// ----- Strings ----- //")
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s[2])                 // strings are an alias for a byte
	fmt.Printf("%v, %T\n", string(s[2]), string(s[2])) // to read as a string

}
