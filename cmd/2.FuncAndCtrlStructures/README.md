
---

### Tutorial 2 README.md (located in `cmd/tutorial_2/`)

```markdown
# Tutorial 2: Functions and Control Structures

This tutorial dives into functions, error handling, and various control structures in Go. It is designed to help you understand how to handle errors gracefully and control your program's flow.

## What You Will Learn

- **Functions:**  
  Define and call functions, pass parameters, and return multiple values.

- **Error Handling:**  
  Learn how Go uses the built-in `error` type. The `intDivision` function returns an error when a division by zero is attempted.

- **Control Structures:**  
  Use `if-else` statements and `switch` statements to manage logic paths. The code demonstrates:
  - An `if-else` block for checking errors and handling division results.
  - A switch statement without an expression to handle multiple conditions.
  - A conditional switch based on the division remainder.

## Code Overview

The main file (`main.go`) does the following:

1. **Greeting:**  
   Calls the `printMe` function to display a simple message.

2. **Division with Error Checking:**  
   Calls `intDivision` with a numerator and a deliberately set denominator of 0 to trigger an error. This demonstrates how to handle errors (e.g., division by zero).

3. **Error Handling with if-else:**  
   Checks if an error occurred. If so, it prints the error message; otherwise, it displays the division result.

4. **Error Handling with Switch:**  
   Provides alternative ways to handle conditions using switch statements.

## How to Run

From the repository root, run:
```bash
go run cmd/2.FuncAndCtrlStructures/main.go
```
