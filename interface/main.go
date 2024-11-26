package main

/*
 * @interface { area(), perimeter() }
 * @types { width, height }
 */
import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return math.Pi * c.radius * c.radius
}

func printShapeData(s shape, i int) {
	fmt.Printf("Index: %v - Area: %.2f - Perimeter: %.2f/n", i, s.area(), s.perimeter())
}

func main() {
	// initilize the instence of shape
	shapes := []shape{
		rect{2, 3},
		circle{2},
	}

	for i, s := range shapes {
		printShapeData(s, i)
	}
}

// :: Outputs
// Index: 0 - Area: 12.00 - Perimeter: 14.00
// Index: 1 - Area: 78.54 - Perimeter: 31.42
