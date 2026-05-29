# Decorator Pattern in Go

This project demonstrates the **Decorator Design Pattern** using a simple pizza example written in Go.

The Decorator Pattern allows behavior to be added to an individual object dynamically without affecting the behavior of other objects from the same class.

## Overview

The example starts with a simple pizza and then adds additional ingredients (decorators) such as:

- Basil (`Margherita`)
- Olives (`Olive`)

Each decorator wraps another `Pizza` instance and extends its behavior by:

- Adding a new description
- Increasing the total cost

## Project Structure

```go
type Pizza interface {
    Cost() float64
    Description() string
}
```

### Base Component

```go
type SimplePizza struct{}
```

A basic mozzarella pizza with:

- Description: `Mozzarella`
- Cost: `10.0`

### Decorator: Margherita

```go
type Margherita struct {
    pizza Pizza
}
```

Adds:

- Basil to the description
- $5.00 to the total cost

### Decorator: Olive

```go
type Olive struct {
    pizza Pizza
}
```

Adds:

- Olives to the description
- $2.00 to the total cost

## How It Works

A decorator wraps another object implementing the same interface.

Example:

```go
pizza := &SimplePizza{}

margherita := &Margherita{
    pizza: pizza,
}
```

Result:

```text
Mozzarella + basil pizza
Price: US$ 15
```

Multiple decorators can be chained together:

```go
pizza := &SimplePizza{}

margheritaWithOlives := &Olive{
    pizza: &Margherita{
        pizza: pizza,
    },
}
```

Result:

```text
Mozzarella + basil + olives pizza
Price: US$ 17
```

## Running the Example

### Prerequisites

- Go 1.20+

### Execute

```bash
go run main.go
```

### Example Output

```text
Mozzarella pizza
Price: US$ 10
-----------------------

Mozzarella + basil pizza
Price: US$ 15
-----------------------

Mozzarella + olives pizza
Price: US$ 12
-----------------------

Mozzarella + basil + olives pizza
Price: US$ 17
-----------------------
```

## Why Use the Decorator Pattern?

Without decorators, adding combinations of ingredients would quickly lead to many classes:

```text
MozzarellaPizza
MozzarellaWithBasilPizza
MozzarellaWithOlivesPizza
MozzarellaWithBasilAndOlivesPizza
...
```

The Decorator Pattern avoids this explosion of types by allowing features to be composed dynamically at runtime.

## Benefits

- Open/Closed Principle compliant
- Flexible composition of behaviors
- Avoids subclass explosion
- Easy to extend with new decorators

## Design Pattern Category

**Structural Pattern**

The Decorator Pattern belongs to the Structural Design Patterns category because it focuses on how objects are composed to form larger structures.

## References

- Gang of Four (GoF) Design Patterns
- https://refactoring.guru/design-patterns/decorator
- Effective Go
