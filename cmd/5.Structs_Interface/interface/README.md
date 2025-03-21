# Interfaces in Go

## What are Interfaces?
Interfaces in Go define a set of method signatures that a struct (or any type) must implement. They provide a way to achieve polymorphism and flexibility in Go without explicitly defining class hierarchies.

## Interface Usage in Code

### 1. **Basic Interface Example**
- `shape` and `measureable` interfaces define `area()` and `perimeter()` methods.
- `geometery` interface embeds both.
- `rectArea` and `circArea` structs implement these interfaces.

```go
type shape interface {
  area() float64
}

type measureable interface {
  perimeter() float64
}

type geometery interface {
  shape
  measureable
}
```

### 2. Implementing the Interface
- `rectArea` and `circArea` define `area()` and `perimeter()` methods, fulfilling the interface contract.

```go
func (r rectArea) area() float64 {
  return r.width * r.height
}

func (c circArea) area() float64 {
  return math.Pi * c.radius * c.radius
}
```

### 3.Error Handling with Interfaces
- Custom error struct implementing `error` interface:

```go
type calculationError struct {
  msg string
}

func (ce calculationError) Error() string {
  return ce.msg
}

```

### 4. Interfaces in Engine Example
- `engine` interface defines milesLeft() uint8
- `gasEngine` and electricEngine structs implement this method.

```go
type engine interface {
  milesLeft() uint8
}

func (e gasEngine) milesLeft() uint8 {
  return e.gallons * e.mph
}

func (e electricEngine) milesLeft() uint8 {
  return e.mpkwh * e.kwh
}

```

### 5. Interface-Based Function
- A function that works with any engine type:

```go
func canMakeIt(e engine, miles uint8) {
  if miles <= e.milesLeft() {
    fmt.Println("You can make it!")
  } else {
    fmt.Println("Need to fuel up first!")
  }
}

```

## Key Takeaways
- Interfaces act as contracts without implementing logic.
- Any struct with the required methods implicitly satisfies an interface.
- Interfaces enable flexible and reusable code.