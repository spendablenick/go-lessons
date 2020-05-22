package main

import "fmt"

/*

if condition {
	// dothis
}

*/

func main() {
	// nicksAge := 15

	// if nicksAge < 13 {
	// 	fmt.Println("Nick is a kid")
	// } else if nicksAge <= 19 {
	// 	fmt.Println("Nick is a teenager")
	// } else {
	// 	fmt.Println("Nick is an adult")
	// }

	// if nicksAge := 20; nicksAge < 13 {
	// 	fmt.Println("Nick is a kid")
	// }

	//switch case

	// switch {
	// case nicksAge < 13:
	// 	fmt.Println("Nick is a kid")
	// 	fallthrough
	// case nicksAge < 19:
	// 	fmt.Println("Nick is a teenager")
	// 	fallthrough
	// default:
	// 	fmt.Println("Nick is ", nicksAge)
	// }

	// switch nicksAge {
	// case 5, 10, 15:
	// 	fmt.Println("Nick is 5, 10 or 15")
	// }

	// fmt.Println("-----------------")

	// if nicksAge < 13 {
	// 	fmt.Println("Nick is less than 13")
	// }
	// if nicksAge < 19 {
	// 	fmt.Println("Nick is a teen")
	// }

	// fmt.Println("Nick is", nicksAge)

	// Control flow directive - for

	// 	for  init | condition | after {
	// // run your code
	// }

	// 1 - exec init
	// 2 - Test condition
	// 3 - Runs you code
	// 4 - executes the after directive
	// 5 - go to step 2

	// for i := 0; i < 5; i++ {
	// 	// fmt.Println("Hello world", i)
	// 	for j := 0; j < 5; j++ {
	// 		fmt.Println("Hello world", i, j)
	// 	}
	// }

	// names := []string{"John", "Jonah", "Steve", "Henry", "Sam", "Paul"}

	// for k := 0; k < len(name); k++ {
	// 	fmt.Println(names[k])
	// }
	// ages := map[string]int{
	// 	"John":  7,
	// 	"Jonah": 12,
	// 	"Sam":   20,
	// }

	// for k, v := range ages {
	// 	fmt.Println(k, v)
	// }

	// name := "Jonah"

	// for k, v := range name {
	// 	fmt.Println(k, string(v))
	// }

	// names := []string{"John", "Jonah", "Steve", "Henry", "Sam", "Paul"}

	// i := 0

	// for {
	// 	if i == len(names) {
	// 		break
	// 	}
	// 	fmt.Println(names[i])
	// 	i++
	// 	if names[i] == "Steve" {
	// 		continue // skip current iteration
	// 	}
	// }

	// writname("Paul", 20)

	//sum, product := addMultiply(5, 4)
	//fmt.Println(sum, product)

	r := 5.0
	c := area(r)
	fmt.Println(c)
}

func area(r float64) float64 {
	calculation := 3.14 * r * r
	return calculation
}

// func writname(name string, age int) {
// 	fmt.Println("Write my name", name, ". I am ", age)

// }

//func addMultiply(a, b int) (sum int, product int) {
// sum = a + b
// product = a * b
//	return sum, product
//}
