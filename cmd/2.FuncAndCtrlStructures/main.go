package main

import (
	"errors"
	"fmt"
)

// main demonstrates various Go concepts such as functions, error handling,
// and control structures (if-else and switch).
func main() {
	// -------------------------------------------------------------------
	// Example 1: Printing a Message
	// -------------------------------------------------------------------
	// Here, we declare a string variable 'printValue' and pass it to our
	// helper function 'printMe' which prints the message.
	var printValue string = "hello world"
	printMe(printValue)

	// -------------------------------------------------------------------
	// Example 2: Integer Division with Error Handling
	// -------------------------------------------------------------------
	// In this example, we perform integer division. We have a numerator and a
	// denominator. Setting the denominator to zero simulates an error case. And to handle that we use error method. We are using error methtod to handle error
	var numerator int = 11
	var denominator int = 0 // Change to a non-zero value (e.g., 2) to see normal behavior

	// intDivision performs the division and returns three values:
	// - The quotient (result of integer division)
	// - The remainder
	// - An error if division by zero occurs
	result, remainder, err := intDivision(numerator, denominator)

	// Check for errors using an if-else control structure.
	// If an error is returned (i.e., err != nil), print the error message.
	// Otherwise, decide what to print based on whether there's a remainder.
	if err != nil {
		// An error occurred during division (for example, dividing by zero).
		fmt.Println("Error:", err.Error())
	} else if remainder == 0 {
		// If there is no remainder, the division was exact.
		fmt.Printf("The result of integer division is %v\n", result)
	} else {
		// If there's a non-zero remainder, display both the quotient and remainder.
		fmt.Printf("The result of integer division is %v with a remainder of %v\n", result, remainder)
	}

	// -------------------------------------------------------------------
	// Example 3: Using a Switch Statement for Flow Control
	// -------------------------------------------------------------------
	// A switch without an expression works like a series of if-else statements.
	// It checks each condition in order until one matches.
	switch {
	case err != nil:
		fmt.Println("Error:", err.Error())
	case remainder == 0:
		fmt.Printf("The result of integer division is %v\n", result)
	default:
		fmt.Printf("The result of integer division is %v with a remainder of %v\n", result, remainder)
	}

	// -------------------------------------------------------------------
	// Example 4: Conditional Switch Based on a Variable
	// -------------------------------------------------------------------
	// In this switch statement, we use the 'remainder' variable directly.
	// This lets us execute different code based on the value of 'remainder'.
	switch remainder {
	case 0:
		fmt.Println("The division was exact.")
	case 1, 2:
		fmt.Println("The division was close.")
	default:
		fmt.Println("The division was not close.")
	}
}

// printMe is a simple helper function that prints the provided message.
// This function demonstrates how to pass a parameter to a function.
func printMe(message string) {
	fmt.Println(message)
}

// intDivision performs integer division and handles division by zero.
// It takes two integer parameters: numerator and denominator.
// It returns three values:
//   - The quotient of the division (integer division)
//   - The remainder of the division
//   - An error if the denominator is zero (to avoid division by zero)
func intDivision(numerator, denominator int) (int, int, error) {
	// Check for division by zero to avoid a runtime error.
	if denominator == 0 {
		// Return zero values for quotient and remainder, along with an error.
		return 0, 0, errors.New("cannot divide by zero")
	}

	// Calculate the quotient and remainder.
	quotient := numerator / denominator
	remainder := numerator % denominator

	// Return the computed values and nil for error (indicating no error occurred).
	return quotient, remainder, nil
}
