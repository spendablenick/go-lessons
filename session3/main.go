package main

import "fmt"

/*

store =
{
	"1234": "Nick",
	"23456": "Bubu"
	"2489423": "james"
}

get key "1234" from Store -> Nick
get key "1234" from Store -> Nick

*/

// func main() {
// 	// var m map[string]int
// 	/*
// 		m = {
// 			"age": 9
// 		}
// 	*/

// 	m := make(map[string][]int)
// 	m["nick"] = []int{90, 45, 80}
// 	m["bubu"] = []int{60, 120, 13}
// 	// m1 := map[string]int{}

// 	// fmt.Println(m)

// 	// delete(m, "bubu")

// 	// _, ok := m["bubu"]

// 	fmt.Println(m)

// }

/*

Operators - set of symbols that tell the compiler to perform a set of actions.

5 types
1. Arithmetic oeprators
2. Assingment Operators


*/

func main() {

	//Arrithmetic operators
	// x := 10
	// y := 17
	// fmt.Println(x + y)
	// fmt.Println(x - y)
	// fmt.Println(x * y)
	// fmt.Println(y / x)
	// fmt.Println(y % x)

	// x++
	// y--

	// Assingment operators
	// var x, y = 10, 17
	// x /= 5 // x = x/5
	// y %= 3 // y = y%3
	// fmt.Println(x, y)

	// Comparison operators
	// var x, y = 10, 10
	// // fmt.Println(x, y, x == y)

	// fmt.Println(x, y, x >= y)
	// fmt.Println(x, y, x <= y)

	// fmt.Println(x, y, x != y)

	// Logic operators
	// var x, y = 10, 17
	// // x /= 5 // x = x/5
	// // y %= 3 // y = y%3
	// fmt.Println(x, y, x > 5 || x < 50)
	// fmt.Println(x, y, x > 5 && x < 50)

	// Bitwise operators
	x := 9  // 0000 1001
	y := 65 // 0100 0001
	// 0100 1001
	// x /= 5 // x = x/5
	// y %= 3 // y = y%3
	fmt.Println(x | y)
	// fmt.Println(x, y, x > 5 && x < 50)
}
