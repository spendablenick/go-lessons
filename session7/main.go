// fmt.Println()
package main

import "fmt"

type now struct {
	day   int
	month int
	year  int
}

func (n now) Year() int {
	return n.year
}

func add1Year(n *now) {
	n.year++
}

func main() {

	now := now{
		day:   9,
		month: 5,
		year:  2020,
	}
	fmt.Println(now.Year())
	add1Year(&now)
	add1Year(&now)
	add1Year(&now)

	fmt.Println(now.Year())
	// now.Year()
}


var x = 5
y = 5
z := 5



