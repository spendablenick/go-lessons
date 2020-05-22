package shape

import (
	"fmt"
)

type Circle struct {
	name   string
	Radius float64
}

func (pass Circle) Area() float64 {

	pass.name = "Big Circle"

	fmt.Println("Hi is am ", pass.name, "and i have a radius of ", pass.Radius)

	return pass.Radius * pass.Radius * 3.14
}
func (pass Circle) Perimeter() float64 {
	return 2 * pass.Radius * 3.14
}

func (pass Circle) Volumn(h float64) float64 {
	return pass.Area() * h
}
