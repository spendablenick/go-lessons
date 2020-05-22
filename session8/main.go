package main

import (
	"fmt"
)

// type Person struct {
// 	Name    string
// 	Age     int
// 	Address [2]string
// 	Height  int
// }

// func (p Person) Speak() {
// 	fmt.Println("Hi my name is ", p.Name)
// }

// func (p *Person) bumpAge() {
// 	p.Age++
// }

// func main() {
// 	var bubu = Person{
// 		Name: "Bubunyo",
// 		Age:  27,
// 		Address: [2]string{
// 			"19 Banana Street",
// 			"East Legon",
// 		},
// 		Height: 180,
// 	}

// 	nick := Person{
// 		"Nick",
// 		27,
// 		[2]string{
// 			"728 East London",
// 			"London",
// 		},
// 		180,
// 	}

// 	fmt.Println(bubu.Age)
// 	bubu.Speak()
// 	bubu.Age = bubu.Age + 1
// 	fmt.Println(bubu.Age)
// 	fmt.Println(nick.Address)
// 	nick.bumpAge()
// 	fmt.Println(nick.Age)

// 	c := shape.Circle{Radius: 30}

// 	fmt.Println("Hi is am ", "and i have a radius of ", c.Radius)

// 	fmt.Println(c.Area())

// }

type Shape interface {
	Area() float64
	Perimeter() float64
}

func SpeakArea(s Shape) {
	fmt.Println("I speak my area as", s.Area())
}

type Rectangle struct {
	Length float64
	Breath float64
}

func (r Rectangle) Area() float64 {
	return r.Breath * r.Length
}
func (r Rectangle) Perimeter() float64 {
	return (r.Breath + r.Length) * 2
}

func (r Rectangle) NumberOfSides() int {
	return 4
}

type Circle struct {
	Radius float64
}

func (r Circle) Area() float64 {
	return r.Radius * r.Radius * 3.14
}
func (r Circle) Perimeter() float64 {
	return 2 * 3.14 * r.Radius
}

func main() {
	rect := Rectangle{Length: 20, Breath: 30}
	circle := Circle{Radius: 20}
	SpeakArea(rect)
	SpeakArea(circle)
}
