package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	showWelcomeMessage()

	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		showMainMenu()
		fmt.Print("Enter your choice (1-6): ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())
		
		switch choice {
		case "1":
			basicPointerDemo(scanner)
		case "2":
			pointerToVariableDemo(scanner)
		case "3":
			slicePointerDemo(scanner)
		case "4":
			arrayFunctionDemo(scanner)
		case "5":
			interactivePointerExercise(scanner)
		case "6":
			fmt.Println("\nThank you for learning about Go pointers! Goodbye!")
			return
		default:
			fmt.Println("\n⚠️ Invalid choice. Please select a number between 1 and 6.")
		}
		
		fmt.Println("\nPress Enter to continue...")
		scanner.Scan()
	}
}

func showWelcomeMessage() {
	fmt.Println("================================================")
	fmt.Println("  WELCOME TO THE INTERACTIVE GO POINTERS TUTORIAL")
	fmt.Println("================================================")
	fmt.Println("This tutorial will help you understand pointers in Go")
	fmt.Println("through interactive examples and exercises.")
	fmt.Println("")
	fmt.Println("What are pointers?")
	fmt.Println("  Pointers are special variables that store memory")
	fmt.Println("  addresses instead of actual values.")
	fmt.Println("")
}

func showMainMenu() {
	fmt.Println("\n-------------------------------------------------")
	fmt.Println("MAIN MENU - Choose a topic to explore:")
	fmt.Println("-------------------------------------------------")
	fmt.Println("1. Basic Pointer Concepts")
	fmt.Println("2. Pointers to Existing Variables")
	fmt.Println("3. Slices and Hidden Pointers")
	fmt.Println("4. Arrays and Pointers with Functions")
	fmt.Println("5. Interactive Exercise")
	fmt.Println("6. Exit Tutorial")
	fmt.Println("-------------------------------------------------")
}

func basicPointerDemo(scanner *bufio.Scanner) {
	clearScreen()
	fmt.Println("=== BASIC POINTER CONCEPTS ===")
	fmt.Println("")
	fmt.Println("In Go, we create pointers using the * symbol:")
	fmt.Println("  var p *int  // Declares a pointer to an integer")
	fmt.Println("")
	
	fmt.Println("Press Enter to see a live example...")
	scanner.Scan()
	
	var p *int // Nil pointer initially
	fmt.Printf("1. We created a pointer: var p *int\n")
	fmt.Printf("   Value of p: %v (this is nil because it doesn't point anywhere yet)\n\n", p)
	
	fmt.Println("Press Enter to initialize the pointer...")
	scanner.Scan()
	
	p = new(int) // Initialize with new()
	fmt.Printf("2. We initialized p with new(int)\n")
	fmt.Printf("   Value of p: %v (this is a memory address now)\n", p)
	fmt.Printf("   Value at this address (*p): %v (default zero value for int)\n\n", *p)
	
	fmt.Println("Press Enter to modify the value...")
	scanner.Scan()
	
	*p = 42 // Modify the value
	fmt.Printf("3. We changed the value with *p = 42\n")
	fmt.Printf("   Value of p: %v (address is unchanged)\n", p)
	fmt.Printf("   Value at this address (*p): %v (now it's 42!)\n\n", *p)
	
	// Interactive part
	fmt.Println("Now it's your turn! Enter a number to store at this pointer address:")
	scanner.Scan()
	userInput := scanner.Text()
	if num, err := strconv.Atoi(userInput); err == nil {
		*p = num
		fmt.Printf("\n✅ Great! You've changed the value to %v\n", *p)
		fmt.Printf("   The pointer p still points to the same address: %v\n", p)
	} else {
		fmt.Println("\n⚠️ That's not a valid number, but that's okay!")
		fmt.Println("   Let's continue with the tutorial.")
	}
	
	fmt.Println("\n📚 KEY TAKEAWAYS:")
	fmt.Println("   1. A pointer stores a memory address")
	fmt.Println("   2. *Type declares a pointer to Type")
	fmt.Println("   3. new() allocates memory and returns a pointer")
	fmt.Println("   4. *p retrieves the value at the address (dereferencing)")
}

func pointerToVariableDemo(scanner *bufio.Scanner) {
	clearScreen()
	fmt.Println("=== POINTERS TO EXISTING VARIABLES ===")
	fmt.Println("")
	fmt.Println("We can create pointers that refer to existing variables:")
	fmt.Println("  var x int = 10")
	fmt.Println("  var p *int = &x  // & gets the address of x")
	fmt.Println("")
	
	fmt.Println("Press Enter to see a live example...")
	scanner.Scan()
	
	var x int = 10
	fmt.Printf("1. We created a regular variable: var x int = 10\n")
	fmt.Printf("   Value of x: %v\n", x)
	fmt.Printf("   Memory address of x: %v (using &x)\n\n", &x)
	
	fmt.Println("Press Enter to create a pointer to x...")
	scanner.Scan()
	
	var p *int = &x
	fmt.Printf("2. We created a pointer p pointing to x: var p *int = &x\n")
	fmt.Printf("   Value of p: %v (this is x's address)\n", p)
	fmt.Printf("   Value at this address (*p): %v (same as x's value)\n\n", *p)
	
	fmt.Println("Press Enter to modify x through the pointer...")
	scanner.Scan()
	
	*p = 20
	fmt.Printf("3. We changed the value with *p = 20\n")
	fmt.Printf("   Value of p: %v (address is unchanged)\n", p)
	fmt.Printf("   Value at this address (*p): %v\n", *p)
	fmt.Printf("   Value of x: %v (x changed too because p points to x!)\n\n", x)
	
	// Interactive part
	fmt.Println("Your turn! Enter a new value for x (using the pointer p):")
	scanner.Scan()
	userInput := scanner.Text()
	if num, err := strconv.Atoi(userInput); err == nil {
		*p = num
		fmt.Printf("\n✅ Great! You changed *p to %v\n", *p)
		fmt.Printf("   And x is now also %v (they share the same memory location)\n", x)
	} else {
		fmt.Println("\n⚠️ That's not a valid number, but that's okay!")
		fmt.Println("   Let's continue with the tutorial.")
	}
	
	fmt.Println("\n📚 KEY TAKEAWAYS:")
	fmt.Println("   1. &variable gets the memory address of a variable")
	fmt.Println("   2. When you modify *p, you're changing the value at the address")
	fmt.Println("   3. If p points to x, then changing *p also changes x")
}

func slicePointerDemo(scanner *bufio.Scanner) {
	clearScreen()
	fmt.Println("=== SLICES AND HIDDEN POINTERS ===")
	fmt.Println("")
	fmt.Println("Slices in Go already contain pointers to underlying arrays.")
	fmt.Println("This is why when you copy a slice, both refer to the same data:")
	fmt.Println("")
	
	fmt.Println("Press Enter to see how slices behave with copying...")
	scanner.Scan()
	
	// Original slice
	originalSlice := []int{10, 20, 30, 40, 50}
	fmt.Printf("1. We created a slice: originalSlice := []int{10, 20, 30, 40, 50}\n")
	fmt.Printf("   originalSlice: %v\n\n", originalSlice)
	
	fmt.Println("Press Enter to create a copy...")
	scanner.Scan()
	
	// Copy the slice
	sliceCopy := originalSlice
	fmt.Printf("2. We copied the slice: sliceCopy := originalSlice\n")
	fmt.Printf("   originalSlice: %v\n", originalSlice)
	fmt.Printf("   sliceCopy: %v\n\n", sliceCopy)
	
	fmt.Println("Press Enter to modify the copy...")
	scanner.Scan()
	
	// Modify the copy
	sliceCopy[2] = 999
	fmt.Printf("3. We modified the copy: sliceCopy[2] = 999\n")
	fmt.Printf("   sliceCopy: %v\n", sliceCopy)
	fmt.Printf("   originalSlice: %v (original changed too!)\n\n", originalSlice)
	
	// Create a new slice with make to demonstrate
	fmt.Println("Press Enter to see how to create a true independent copy...")
	scanner.Scan()
	
	// Create a true copy
	trueSliceCopy := make([]int, len(originalSlice))
	copy(trueSliceCopy, originalSlice)
	
	fmt.Printf("4. We created a true copy with make and copy function:\n")
	fmt.Printf("   trueSliceCopy := make([]int, len(originalSlice))\n")
	fmt.Printf("   copy(trueSliceCopy, originalSlice)\n\n")
	
	// Modify the true copy
	fmt.Println("Press Enter to modify this true copy...")
	scanner.Scan()
	
	trueSliceCopy[2] = 777
	fmt.Printf("5. We modified the true copy: trueSliceCopy[2] = 777\n")
	fmt.Printf("   trueSliceCopy: %v\n", trueSliceCopy)
	fmt.Printf("   originalSlice: %v (original unchanged!)\n\n", originalSlice)
	
	// Interactive part
	fmt.Println("Your turn! Enter an index (0-4) to modify in the originalSlice:")
	scanner.Scan()
	idxStr := scanner.Text()
	
	if idx, err := strconv.Atoi(idxStr); err == nil && idx >= 0 && idx < len(originalSlice) {
		fmt.Printf("Now enter a new value for index %d:\n", idx)
		scanner.Scan()
		valStr := scanner.Text()
		
		if val, err := strconv.Atoi(valStr); err == nil {
			originalSlice[idx] = val
			fmt.Printf("\n✅ You changed originalSlice[%d] to %d\n", idx, val)
			fmt.Printf("   originalSlice: %v\n", originalSlice)
			fmt.Printf("   sliceCopy: %v (also changed!)\n", sliceCopy)
			fmt.Printf("   trueSliceCopy: %v (unchanged!)\n", trueSliceCopy)
		} else {
			fmt.Println("\n⚠️ That's not a valid number")
		}
	} else {
		fmt.Println("\n⚠️ That's not a valid index")
	}
	
	fmt.Println("\n📚 KEY TAKEAWAYS:")
	fmt.Println("   1. Slices in Go contain pointers to underlying arrays")
	fmt.Println("   2. When you copy a slice, both copies refer to the same data")
	fmt.Println("   3. To create an independent copy, use make() and copy()")
}

func arrayFunctionDemo(scanner *bufio.Scanner) {
	clearScreen()
	fmt.Println("=== ARRAYS AND POINTERS WITH FUNCTIONS ===")
	fmt.Println("")
	fmt.Println("When we pass arrays to functions in Go, they are copied by default.")
	fmt.Println("This is inefficient for large arrays. Using pointers solves this:")
	fmt.Println("")
	
	// Create and display original array
	originalArray := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("1. We created an array: originalArray := [5]int{10, 20, 30, 40, 50}\n")
	fmt.Printf("   originalArray: %v\n", originalArray)
	fmt.Printf("   Memory address: %p\n\n", &originalArray)
	
	fmt.Println("Press Enter to see what happens with copy-by-value...")
	scanner.Scan()
	
	// Demonstrate copy by value
	modifiedByValue := modifyArrayByValue(originalArray)
	fmt.Printf("2. We called modifyArrayByValue(originalArray):\n")
	fmt.Printf("   Return value: %v\n", modifiedByValue)
	fmt.Printf("   originalArray: %v (unchanged!)\n\n", originalArray)
	
	fmt.Println("Press Enter to see what happens with pointers...")
	scanner.Scan()
	
	// Demonstrate using pointers
	modifiedByPointer := modifyArrayByPointer(&originalArray)
	fmt.Printf("3. We called modifyArrayByPointer(&originalArray):\n")
	fmt.Printf("   Return value: %v\n", modifiedByPointer)
	fmt.Printf("   originalArray: %v (modified!)\n\n", originalArray)
	
	// Reset array for interactive part
	originalArray = [5]int{10, 20, 30, 40, 50}
	
	// Interactive part
	fmt.Println("Let's try one more time interactively.")
	fmt.Printf("Current array: %v\n\n", originalArray)
	
	fmt.Println("Choose how to pass the array to the function:")
	fmt.Println("1. By value (copy)")
	fmt.Println("2. By pointer (reference)")
	fmt.Print("Enter choice (1 or 2): ")
	
	scanner.Scan()
	choice := scanner.Text()
	
	if choice == "1" {
		result := modifyArrayByValue(originalArray)
		fmt.Printf("\n✅ Function returned: %v\n", result)
		fmt.Printf("   Original array after function call: %v\n", originalArray)
		fmt.Println("   The original array is unchanged because we passed a copy!")
	} else if choice == "2" {
		result := modifyArrayByPointer(&originalArray)
		fmt.Printf("\n✅ Function returned: %v\n", result)
		fmt.Printf("   Original array after function call: %v\n", originalArray)
		fmt.Println("   The original array changed because we passed a pointer!")
	} else {
		fmt.Println("\n⚠️ Invalid choice, but that's okay!")
	}
	
	fmt.Println("\n📚 KEY TAKEAWAYS:")
	fmt.Println("   1. Arrays are copied when passed to functions by default")
	fmt.Println("   2. Use pointers to avoid copying large arrays (more efficient)")
	fmt.Println("   3. When using pointers, changes in the function affect the original array")
	fmt.Println("   4. This is especially important for large data structures")
}

func interactivePointerExercise(scanner *bufio.Scanner) {
	clearScreen()
	fmt.Println("=== INTERACTIVE POINTER EXERCISE ===")
	fmt.Println("")
	fmt.Println("Let's test your understanding of pointers!")
	fmt.Println("")
	
	var score int = 0
	
	// Question 1
	fmt.Println("Question 1: What symbol is used to declare a pointer type in Go?")
	fmt.Println("a) &")
	fmt.Println("b) *")
	fmt.Println("c) #")
	fmt.Println("d) @")
	fmt.Print("Your answer: ")
	
	scanner.Scan()
	if strings.ToLower(scanner.Text()) == "b" {
		fmt.Println("✅ Correct! * is used to declare a pointer type.")
		score++
	} else {
		fmt.Println("❌ Incorrect. * is used to declare a pointer type.")
	}
	
	fmt.Println("\nPress Enter to continue...")
	scanner.Scan()
	
	// Question 2
	fmt.Println("\nQuestion 2: What symbol is used to get the address of a variable?")
	fmt.Println("a) &")
	fmt.Println("b) *")
	fmt.Println("c) ->")
	fmt.Println("d) @")
	fmt.Print("Your answer: ")
	
	scanner.Scan()
	if strings.ToLower(scanner.Text()) == "a" {
		fmt.Println("✅ Correct! & is used to get the address of a variable.")
		score++
	} else {
		fmt.Println("❌ Incorrect. & is used to get the address of a variable.")
	}
	
	fmt.Println("\nPress Enter to continue...")
	scanner.Scan()
	
	// Question 3
	fmt.Println("\nQuestion 3: What happens when you modify a slice that was copied from another slice?")
	fmt.Println("a) Nothing happens to the original slice")
	fmt.Println("b) The original slice is also modified")
	fmt.Println("c) The program crashes")
	fmt.Println("d) The copy is automatically deleted")
	fmt.Print("Your answer: ")
	
	scanner.Scan()
	if strings.ToLower(scanner.Text()) == "b" {
		fmt.Println("✅ Correct! The original slice is also modified because slices contain pointers to the underlying array.")
		score++
	} else {
		fmt.Println("❌ Incorrect. The original slice is also modified because slices contain pointers to the underlying array.")
	}
	
	fmt.Println("\nPress Enter to continue...")
	scanner.Scan()
	
	// Question 4
	fmt.Println("\nQuestion 4: What's the primary benefit of using pointers with function parameters?")
	fmt.Println("a) It makes the code more complicated")
	fmt.Println("b) It allows the function to modify the original data")
	fmt.Println("c) It prevents any data modification")
	fmt.Println("d) It makes the program run slower")
	fmt.Print("Your answer: ")
	
	scanner.Scan()
	if strings.ToLower(scanner.Text()) == "b" {
		fmt.Println("✅ Correct! Using pointers allows functions to modify the original data and avoid unnecessary copying.")
		score++
	} else {
		fmt.Println("❌ Incorrect. Using pointers allows functions to modify the original data and avoid unnecessary copying.")
	}
	
	// Show final score
	fmt.Printf("\nYour final score: %d/4\n", score)
	if score == 4 {
		fmt.Println("🎉 Perfect! You've mastered the basics of Go pointers!")
	} else if score >= 2 {
		fmt.Println("👍 Good job! You're getting the hang of pointers.")
	} else {
		fmt.Println("📚 Keep learning! Pointers can be tricky but they're very useful.")
	}
	
	fmt.Println("\nFeel free to revisit any section from the main menu to strengthen your understanding.")
}

// Helper functions for the demos

func modifyArrayByValue(arr [5]int) [5]int {
	fmt.Printf("   Inside function (by value), array address: %p\n", &arr)
	for i := range arr {
		arr[i] *= 2
	}
	return arr
}

func modifyArrayByPointer(arr *[5]int) [5]int {
	fmt.Printf("   Inside function (by pointer), array address: %p\n", arr)
	for i := range arr {
		arr[i] *= 2
	}
	return *arr
}

func clearScreen() {
	fmt.Print("\033[H\033[2J") // ANSI escape code to clear screen
}
