package main

import (
	"fmt"
	"time"
)

// ==============================================
//  Golang Tutorial: Arrays, Slices, Maps & Loops 
// ==============================================
// Go’s core data structures:
// 
// ARRAYS:
// - Fixed-length collections where all elements are of the same type.
// - Stored in contiguous memory: for example, int32 values are 4 bytes each,
//   so elements are located at addresses x, x+4, x+8, etc.
// - Benefits: Faster CPU access via pointer arithmetic and efficient caching.
// - Arrays don’t store extra metadata (like length or capacity) unlike slices.
// - When declared, elements get default values (0 for int32).
//
// SLICES:
// - Dynamic wrappers around arrays.
// - Can grow using the built-in append() function.
// - When appending, if the underlying array lacks space, a new array is created
//   (with extra capacity) and values are copied over.
// - Preallocating capacity (using make) is great for performance if you know the size ahead.
//
// MAPS:
// - Key-value pairs for fast lookup.
// - Accessing a non-existent key returns the default value (e.g., 0 for uint8).
// - The second boolean value in a lookup lets you know if the key exists.
//
// LOOPS:
// - Use `range` to iterate over arrays, slices, and maps.
// - Go doesn’t have a built-in while loop; you can simulate it with for loops.
//
// BONUS: A speed test to see how preallocating capacity improves performance.
// ==============================================

func main() {
	// ---------------------------
	//   ARRAYS in Golang
	// ---------------------------
	// Arrays are fixed in length and store data contiguously.
	// This allows for fast, cache-friendly access.
	var intArr [3]int32 // Fixed length of 3; default value: [0, 0, 0]
	fmt.Println("Initial intArr[0]:", intArr[0])
	fmt.Println("Slice of intArr (indices 1 to 3):", intArr[1:3]) // Output: [0 0]

	// Modify the array by setting index 1 to 123.
	intArr[1] = 123
	fmt.Println("intArr after modification:", intArr[0:3]) // Output: [0 123 0]

	// Print memory locations for each element.
	// Since int32 is 4 bytes, each address is 4 bytes apart.
	fmt.Println("\nMemory Locations:")
	fmt.Println("Address of intArr[0]:", &intArr[0])
	fmt.Println("Address of intArr[1]:", &intArr[1])
	fmt.Println("Address of intArr[2]:", &intArr[2])
	// This shows the contiguous allocation: each subsequent element is 4 bytes ahead.

	// ---------------------------
	//   Array Initialization Variants
	// ---------------------------
	fmt.Println("\nAnother way to initialize arrays:")
	var scndArr [4]int32 = [4]int32{1, 2, 3, 4}
	fmt.Println("scndArr:", scndArr)

	fmt.Println("\nInitialize array using shorthand (:=):")
	thrdArr := [5]int32{1, 2, 3, 4, 5}
	fmt.Println("thrdArr:", thrdArr)

	fmt.Println("\nInitialize array using [...] to infer length:")
	frthArr := [...]int32{1, 2, 3, 4, 5, 6}
	fmt.Println("frthArr:", frthArr)

	// ---------------------------
	//   SLICES in Golang
	// ---------------------------
	// Slices are dynamic and built on top of arrays.
	// They allow for flexibility with the append() function.
	fmt.Println("\nSlices:")
	var intSlice []int32 = []int32{4, 5, 6} // Slice literal (no length specified)
	fmt.Println("Initial intSlice:", intSlice)

	// Append a new element (7) to intSlice.
	intSlice = append(intSlice, 7)
	fmt.Println("After appending 7:", intSlice)

	fmt.Println("\nAnother way to initialize a slice using shorthand (:=):")
	intSlice2 := []int32{1, 2, 3}
	fmt.Println("Initialized intSlice2:", intSlice2)

	fmt.Println("\nAppending intSlice2 to intSlice using spread operator:")
	intSlice = append(intSlice, intSlice2...)
	fmt.Println("Combined slice:", intSlice)

	// Demonstrate what happens when appending exceeds capacity.
	var appendSlice []int32 = []int32{1, 2, 3}
	fmt.Printf("\nBefore appending: Length = %v, Capacity = %v\n", len(appendSlice), cap(appendSlice))
	fmt.Println("Appending 4 into appendSlice...")
	appendSlice = append(appendSlice, 4)
	fmt.Printf("After appending: Length = %v, Capacity = %v\n", len(appendSlice), cap(appendSlice))
	// Under the hood, if there isn’t enough room, Go allocates a new array with greater capacity.

	// Creating slices using the make() function.
	fmt.Println("\nCreating slice using make() with length 3:")
	var intSlice3 []int32 = make([]int32, 3)
	fmt.Println("intSlice3:", intSlice3)

	// Creating a slice with a defined length and capacity.
	var intSlice4 []int32 = make([]int32, 4, 5) // Length = 4, Capacity = 5
	fmt.Println("intSlice4 (length 4, capacity 5):", intSlice4)
	// Preallocating capacity is beneficial when you expect many appends.

	// ---------------------------
	//   MAPS in Golang
	// ---------------------------
	// Maps store key-value pairs.
	// If you access a non-existent key, you get the default value.
	// To check if a key exists, use the second value returned.
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println("\nEmpty map myMap:", myMap)

	myMap2 := map[string]uint8{"Adam": 23, "Sarah": 24, "Jadon": 26}
	fmt.Println("Accessing Sarah's age:", myMap2["Sarah"])

	// Accessing a non-existent key "Jason" returns 0 (default for uint8).
	fmt.Println("Accessing non-existent key 'Jason':", myMap2["Jason"])

	// Check if key exists using the second return value.
	age, exists := myMap2["Adam"]
	if exists {
		fmt.Printf("Found Adam with age: %v\n", age)
	} else {
		fmt.Println("Invalid Name")
	}

	// Delete a key from the map.
	delete(myMap2, "Adam")

	// ---------------------------
	//   LOOPS in Golang
	// ---------------------------
	// Loops let you iterate over arrays, slices, and maps using `range`.
	fmt.Println("\nIterating over map myMap2 (keys only):")
	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}

	fmt.Println("\nIterating over map myMap2 (keys and values):")
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	fmt.Println("\nIterating over array intArr (index and value):")
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	// Simulate a while loop using for.
	fmt.Println("\nSimulating a while loop:")
	var p int
	for p < 10 {
		fmt.Println("p =", p)
		p += 2
	}

	// Another loop example with initialization, condition, and post statement.
	var i, q int = 2, 0
	for q < 10 {
		fmt.Println("Loop value q =", q)
		q = i + q
	}

	// Run the speed test to see how preallocating capacity improves performance.
	speedTest()
}

// -----------------------------------------------------------------------------
// SPEED TEST: Demonstrating the Benefits of Preallocating Slice Capacity
// -----------------------------------------------------------------------------
// We append 'n' elements to two slices:
// 1. testSlice: Created empty with no preallocation.
// 2. testSlice2: Preallocated with capacity 'n'.
func speedTest() {
	var n int = 10000 // Number of iterations (adjusted for demo purposes)
	var testSlice = []int{}          // No preallocation; capacity starts at 0.
	var testSlice2 = make([]int, 0, n) // Preallocated with capacity 'n'.

	fmt.Println("\nSpeed Test:")
	fmt.Printf("Time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Time with preallocation: %v\n", timeLoop(testSlice2, n))
}

// timeLoop measures how long it takes to append n elements to a slice.
func timeLoop(slice []int, n int) time.Duration {
	startTime := time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(startTime)
}
