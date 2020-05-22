package main

import "fmt"

// const (
// 	DAY  = "Saturday"
// 	WEEK = 30
// 	YEAR = 2020
// )

// func main() {
// 	fmt.Println(DAY, WEEK, YEAR)
// }

//package 1
// var (
// 	FirstName string
// 	firstName = "John"
// )

// func main() {
// 	firstName = "Nick"
// 	// FirstName = "Bubu"
// 	fmt.Println("Hi my name is ", firstName, "His name is ", FirstName)
// }

// //package 2
// var firstame string

// func main() {
// 	firstName = "Nick"
// 	fmt.Println("Hi my name is ", name)
// }

// Arrays
/*
34,54,65,34
*/
// func main() {
// 	// var scores [5]int
// 	// fmt.Println(scores[0])
// 	// scores[2] = 5
// 	// fmt.Println(scores)

// 	names := [4]string{"john", "peter", "pete", "jim"}

// 	names[3] = "jones"

// 	fmt.Println(names)
// 	// scores := [...]int{34, 54, 65, 34, 60, 45}
// 	// // fmt.Println(scores[0])
// 	// // scores[2] = 5
// 	// fmt.Println(scores)

// }

// Slices
/*
34,54,65,34
*/
func main() {
	// var scores [5]int
	// fmt.Println(scores[0])
	// scores[2] = 5
	// fmt.Println(scores)

	// namesArr := [4]string{"john", "peter", "pete", "jim"}

	// two := [5]string{}
	// nameSlc := []string{}

	nameSlc := make([]int, 4, 10)

	nameSlc[3] = 7

	fmt.Println(len(nameSlc), cap(nameSlc))

	nameSlc = append(nameSlc, 3, 8, 5, 0, 23, 78, 55)
	fmt.Println(nameSlc, len(nameSlc), cap(nameSlc))

	nameSlc = append(nameSlc, 3, 8, 5, 0, 23, 78, 55)
	fmt.Println(nameSlc, len(nameSlc), cap(nameSlc))

	nameSlc = append(nameSlc, 3, 8, 5, 0, 23, 78, 55)
	fmt.Println(nameSlc, len(nameSlc), cap(nameSlc))

	// scores := [...]int{34, 54, 65, 34, 60, 45}
	// // fmt.Println(scores[0])
	// // scores[2] = 5
	// fmt.Println(scores)
	/*

		6 bytes
		arr ->	[ | | | | | | | | | ]
				 1 2 3 4 5 6 7 8 9 10

		arr2 ->	[ | | | | | | | | |  |  ]
			     1 2 3 4 5 6 7 8 9 10 11

		scl := make([]int, 6, 10)

		slc =
		pointer to the first element  = arr2-1,
		length of slice = 11,
		capacity of the slice = 11
		len(slc) = 6
	*/

	// arr := [6]int{}

}
