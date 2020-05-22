package main

import "fmt"

// /*
//  n = 6

//  1+2+3+4+5+6
// fib
//  n = 10
//  1+2+3+4+5........... +99+100

// */

// func main() {
// 	//	fmt.Println("Hello world")

// 	fib(6)

// }

// func fib(n int) {
// 	x := 0
// 	for i := 0; i <= n; i++ {
// 		x = x + i
// 	}

// 	fmt.Println(x)
// }

/*
 */

// func main() {
// 	x := 5
// 	fib := func(n int) int {
// 		x = 0
// 		for i := 0; i <= n; i++ {
// 			x = x + i
// 		}
// 		return x
// 	}
// 	fmt.Println(fib(6))
// }

// /*
// */
// func main(){

// }

// var prod int = 10

// func main() {
// 	x := 0
// 	for i := 0; i <= 12; i++ {
// 		sum := x + i
// 		fmt.Println(sum)
// 	}
// 	fmt.Println(sum, prod)

// 	fmt.Println(x)
// }

/*
 */
// func main() {

// 	clacArea := func(pi float64) func(float64) {
// 		pi = pi * 2
// 		fmt.Println("setting  pi", pi)
// 		return func(r float64) {
// 			area := pi * r * r
// 			fmt.Println("Using pi ...")
// 			fmt.Println(area)
// 		}
// 	}

// 	cpi := clacArea(3.14)
// 	cpi(5)
// 	cpi(10)

// 	clacArea(3.14)(5)

// 	cpi1 := clacArea(1.57)
// 	cpi1(5)

// }

// func main() {

// 	clacBottomArrea := func(r float64) func(float64) {
// 		bottomArea := 3.14 * r * r
// 		return func(height float64) {
// 			volume := bottomArea * height
// 			fmt.Println("The volume of tank with diameter ", r*2, "is", volume)
// 		}
// 	}

// 	tank1 := clacBottomArrea(5)
// 	tank1(3)  // when tank is 3 meters
// 	tank1(5)  // when tank is 5 meters
// 	tank1(10) // when tank is 10 meters

// 	tank2 := clacBottomArrea(2.5)
// 	tank2(3)  // when tank is 3 meters
// 	tank2(5)  // when tank is 5 meters
// 	tank2(10) // when tank is 10 meters

// }

// /**/
// func main() {
// 	fmt.Println(addAllNumber(1, 5, 4))

// }

// func addAllNumber(n ...int) int {
// 	sum := 0
// 	for _, v := range n {
// 		sum += v
// 	}
// 	return sum
// }

// /**/
// func main() {
// 	addAllNumber([]int{1, 4, 5, 24}, func(sum int) {
// 		fmt.Println(sum)
// 	})
// }

// func addAllNumber(n []int, callback func(int)) {
// 	sum := 0
// 	for _, v := range n {
// 		sum += v
// 	}
// 	callback(sum)
// }

// /**/
func main() {
	fmt.Println(factorial(6))
	// 1*2*3*4*5*6 = 720
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n + factorial(n-1)
}

/*
fact(6) = 6 + fact(5)
fact(5) = 5 + fact(4)
fact(4) = 4 + fact(3)
fact(3) = 3 + fact(2)
fact(2) = 2 + fact(1)
fact(1) = 1
*/
