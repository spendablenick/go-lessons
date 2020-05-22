package main

import (
	"fmt"
)

func main() {

	//assignment2()
	assignment4()

}

func assignment2() {
	//Assignment 2

	const x = 5
	const y = 10

	grid := [y][x]int{}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i%2 == 0 && j%2 == 0 {
				grid[i][j] = 1
			} else if i%2 == 1 && j%2 == 1 {
				grid[i][j] = 1
			} else {
				grid[i][j] = 0
			}
		}
	}
	for k := 0; k < y; k++ {
		fmt.Println(grid[k])
	}
}

func assignment3() {
	//Assignment 3

	thingsLost := []string{"nail", "shoe", "horse", "rider", "house", "message", "battle", "kingdom"}
	for k := 0; k < len(thingsLost)-1; k++ {
		fmt.Println("For want of a", thingsLost[k], "the", thingsLost[k+1], "was lost.")
	}

	// fmt.Println("And all for the want of a", thingsLost[0], "\b.")
	fmt.Printf("And all for the want of a %s.\n", thingsLost[0])

	/// do you see this
}

func assignment4() {
	//Assignment 4

	word := "animal"
	var points int
	letters := map[string]int{
		"a": 1, "e": 1, "i": 1, "o": 1, "u": 1, "l": 1, "n": 1, "r": 1, "s": 1, "t": 1,
		"d": 2, "g": 2,
		"b": 3, "c": 3, "m": 3, "p": 3,
		"f": 4, "h": 4, "v": 4, "w": 4, "y": 4,
		"k": 5,
		"j": 8, "x": 8,
		"q": 10, "z": 10,
		//Is there a way of grouping all letters with the same points value?
	}

	wordLength := len(word)
	for l := 0; l < wordLength; l++ {
		points = points + letters[string(word[l])]
	}
	fmt.Println("The score for", word, "is", points)
}

func assignment5() {
	//Assignment 5

	//array := [4][2]int{1: {20, 21}, 3: {40, 41}}

	// //
	// //this came

	// arr := []interface{}{}

	// list := [5]int{1, 2, 3, 4, 5}
	fmt.Println(list)

}
