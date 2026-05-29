package main

import (
	"fmt"
)

// Interface
type Pizza interface {
	Cost() float64
	Description() string
}

// Base component
type SimplePizza struct{}

func (sp *SimplePizza) Cost() float64 {
	return 10.0
}

func (sp *SimplePizza) Description() string {
	return "Mozzarella"
}

// Decorator 1: Margherita
type Margherita struct {
	pizza Pizza
}

func (m *Margherita) Description() string {
	return m.pizza.Description() + " + basil"
}

func (m *Margherita) Cost() float64 {
	return m.pizza.Cost() + 5.0
}

// Decorator 2: Olive
type Olive struct {
	pizza Pizza
}

func (o *Olive) Description() string {
	return o.pizza.Description() + " + olives"
}

func (o *Olive) Cost() float64 {
	return o.pizza.Cost() + 2.0
}

func Print(p Pizza) {
	fmt.Printf("%s pizza\n", p.Description())
	fmt.Printf("Price: US$ %v\n", p.Cost())
	fmt.Println("-----------------------")
}

func main() {
	// Create a simple pizza
	pizza := &SimplePizza{}
	Print(pizza)

	// Decorate with basil
	margherita := &Margherita{pizza: pizza}
	Print(margherita)

	// Decorate with olives
	olive := &Olive{pizza: pizza}
	Print(olive)

	// Decorate with basil and olives
	margheritaWithOlives := &Olive{pizza: &Margherita{pizza: pizza}}
	Print(margheritaWithOlives)

}
