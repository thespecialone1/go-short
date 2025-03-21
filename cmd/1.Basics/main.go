package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Print a greeting.
	fmt.Println("hello world")

	// -------------------------------
	// Variables and Types
	// -------------------------------

	// Declare an integer variable.
	// Go offers various integer sizes (int8, int16, int32, int64).
	var varInt int16 = 12

	// Unsigned integers (e.g., uint16) store only non-negative values,
	// allowing for a larger maximum value in the same memory.
	var intNum uint16
	fmt.Println(varInt)
	fmt.Println(intNum)

	// Floating-point numbers:
	// Go supports float32 and float64 (there is no generic "float" type).
	var floatNum float32               // Defaults to 0.0
	var float32Num float32 = 123456.9567 // May round due to 32-bit precision
	var float64Num float64 = 123456.9567 // More precise due to 64-bit precision
	fmt.Println(floatNum)
	fmt.Println(float32Num)
	fmt.Println(float64Num)

	// -------------------------------
	// Strings and Characters
	// -------------------------------

	// Concatenating strings with the + operator.
	var myString string = "hello" + "world"
	fmt.Println(myString)

	// The len() function returns the number of bytes, not runes (Unicode characters).
	fmt.Println(len("test")) // Output: 4

	// Characters outside the basic ASCII set may use more than one byte.
	fmt.Println(len("A"))   // Output: 1
	fmt.Println(len("γ"))   // Output: 2

	// To count actual Unicode characters (runes), use utf8.RuneCountInString.
	fmt.Println(utf8.RuneCountInString("γ")) // Output: 1

	// Runes represent Unicode code points.
	// Defined with single quotes, they print as their numeric code.
	var myRune rune = 'a'
	fmt.Println(myRune) // Output: 97

	// -------------------------------
	// Booleans
	// -------------------------------

	// Boolean values: true or false.
	var myBoolean bool = false
	fmt.Println(myBoolean) // Output: false

	// Default values:
	// - Numeric types and runes default to 0.
	// - Strings default to "".
	// - Booleans default to false.

	// -------------------------------
	// Short Variable Declarations
	// -------------------------------

	// You can let Go infer the type.
	var myVar = "text" // Equivalent to: var myVar string = "text"
	myVaar := "text"   // Shorthand declaration (:=)
	fmt.Println(myVar)
	fmt.Println(myVaar)

	// Declaring multiple variables on one line.
	var1, var2, var3 := 1, 2, 3
	fmt.Println(var1, var2, var3) // Output: 1 2 3

	// Best Practice:
	// Use explicit types when the type isn't obvious or when clarity is required,
	// such as when assigning the return value of a function.

	// -------------------------------
	// Constants
	// -------------------------------

	// Constants cannot be changed once defined.
	// They must be declared and initialized at the same time.
	const myConst string = "const_value"
	fmt.Println(myConst)

	// Example: using a constant for a well-known value.
	const pi float32 = 3.1415
	fmt.Println(pi)
}
