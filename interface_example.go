package main

import "math" //შემოიტანს 3.14-ს

type shape interface {
	Area() float64
	Perimeter() float64
}

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type square struct {
	side float64
}

func (sq square) Area() float64 {
	return sq.side * sq.side
}

func (sq square) Perimeter() float64 {
	return sq.side * 4
}

func (sq square) calculateDiagonal() float64 {
	return math.Sqrt(sq.side * sq.side * 2)
}
