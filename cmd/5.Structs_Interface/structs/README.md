# Structs in Go

## What are Structs?
Structs are composite data types that group together variables under a single type, allowing you to define complex objects.

## Struct Usage in Code

### 1. **Defining a Struct**
-   A `budgetExpense` struct with multiple fields, including another struct (`owner`).

```go
type budgetExpense struct {
  income   uint8
  salary   uint8
  expense  uint8
  ownerInfo owner
}

type owner struct {
  name string
}
```

### 2. Creating and Initializing Structs
- Using struct literals:

```go
var myBudget = budgetExpense{income: 55, salary: 50, ownerInfo: owner{"Alex"}, expense: 10}
```

- Anonymous struct:

```go
var anonStruct = struct {
  income uint
  salary uint
}{25, 30}
```

### 3. Methods in Structs
- Methods allow struct instances to have behavior:

```go
func (b budgetExpense) profit() uint8 {
  return b.income - b.expense
}
```

### 4. Comparing Struct Instances
- Function to compare two budgetExpense instances:

```go
func whoMadeMore(b1, b2 budgetExpense) {
  if b1.profit() > b2.profit() {
    fmt.Printf("%v made more profit\n", b1.ownerInfo.name)
  } else {
    fmt.Printf("%v made more profit\n", b2.ownerInfo.name)
  }
}

```

## Key Takeaways
- Structs group related data into a single type.
- Structs can contain fields of different types, including other structs.
- Methods can be associated with structs for encapsulated behavior.
