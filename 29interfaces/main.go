package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}
type Square struct {
	Length float64
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

// it will work fine even if yyou don't use * here and pass reference
// but when you're working with slices and all its good to pass reference

func (c *Circle) Area() float64 {
	return math.Pi * (c.radius * c.radius)
}
func (s *Square) Area() float64 {
	return s.Length * s.Length
}

func (r *Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

func main() {
	fmt.Println("testing interfaces...")

	cir := Circle{
		radius: 3,
	}
	rect := Rectangle{
		Length:  10,
		Breadth: 20,
	}
	sqr := Square{
		Length: 20,
	}
	objs := []Shape{&rect, &sqr, &cir}

	for _, obj := range objs {
		fmt.Printf("area of %T is value:%v \n", obj, GetArea(obj))
	}
}

func GetArea(obj Shape) float64 {
	return obj.Area()
}
