package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	peri() float64
}

type Rect struct {
	width  float64
	height float64
}

type Cir struct {
	rad float64
}

func (obj Rect) area() float64 {
	return obj.height * obj.width
}

func (obj Rect) peri() float64 {
	return 2 * (obj.height + obj.width)
}

func (obj Cir) area() float64 {
	return math.Pi * obj.rad * obj.rad
}

func (obj Cir) peri() float64 {
	return 2 * math.Pi * obj.rad
}

// Function receiving parameter as empty interface
func empty_interface(i interface{}) {
	fmt.Printf("%T --> %v\n", i, i)
}

func main() {
	fmt.Println("Welcome to go-interface")

	var a Shape
	a = Rect{4.0, 5.0}
	b := Rect{4.0, 5.0}
	fmt.Printf("Type of variable a is : %T\n", a)
	fmt.Printf("Value of variable a is : %v\n", a)

	fmt.Printf("Area of rectangle is %f\n", a.area())
	fmt.Printf("Perimeter of rectangle is %f\n", a.peri())
	fmt.Println("whether a == b or not ", a == b)

	a = Cir{3.0}
	fmt.Printf("\nType of variable a is : %T\n", a)
	fmt.Printf("Value of variable a is : %v\n", a)

	fmt.Printf("Area of circle is %f\n", a.area())
	fmt.Printf("Perimeter of circle is %f\n", a.peri())

	r := Cir{2}
	fmt.Printf("\nType of variable a is : %T\n", r)
	fmt.Printf("Value of variable a is : %v\n", r)

	fmt.Printf("Area of circle is %f\n", r.area())
	fmt.Printf("Perimeter of circle is %f\n", r.peri())

	empty_interface(r)
	empty_interface(b)
	empty_interface(2)
}
