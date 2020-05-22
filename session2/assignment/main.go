package main

import "fmt"

func main() {
	var array [6][6]int

	array[1][1] = 1
	array[2][2] = 2
	array[3][3] = 3
	array[4][4] = 4
	array[5][5] = 5

	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println(array[2])
	fmt.Println(array[3])
	fmt.Println(array[4])
	fmt.Println(array[5])

}
