package main

import "fmt"

// func main() {
// 	// v := "a⌘c"
// 	// // fmt.Println(v, len(v))
// 	// for i := 0; i < len(v); i++ {
// 	// 	fmt.Println(string(v[i]))
// 	// }

// 	// r := 'a'
// 	// arr := []rune{'h', 'e', 'l', 'l', 'o'}
// 	// fmt.Println(r, arr, string(arr))

// 	fmt.Println(len("世界"), utf8.RuneCountInString("世界"))
// }

/*

10101010
2^*

rune - single point of code representing a character

a⌘c
'a', '⌘', 'c'

byte, byte, byte, byte

*/

func main() {
	// name := "bubu"

	// pName := &name

	// fmt.Println(name, *pName)

	v := "Nick"
	som(&v)
	fmt.Println(v)

}

func som(v *string) {
	*v = "bubu"
	fmt.Println(*v)
}
