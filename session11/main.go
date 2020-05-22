package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// defer

func Copy(src, dest string) (int64, error) {

	srcFile, err := os.Open(src)
	if err != nil {
		log.Fatal("Src File counld not be opened", err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		log.Fatal("Dest File counld not be open", err)
	}
	defer destFile.Close()

	return io.Copy(destFile, srcFile)

}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	fmt.Println(i)
	return
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

func c() (i int) {
	defer func() {
		i++
	}()
	return 1
}

func d() (i int) {
	arr := []string{"Bubu", "Nick", "John"}
	for i := 0; i <= len(arr); i++ {
		if i >= len(arr) {
			panic("Dude, this index for this array is illegal")
		}
		fmt.Println(i, " - ", arr[i])
	}
	return len(arr)
}

func main() {
	// defer fmt.Println("First")
	// fmt.Println("Second")
	// defer fmt.Println("Third")

	// // LIFO - Last In First Out

	// w, err := Copy("file.txt", "copy.txt")
	// if err != nil {
	// 	log.Fatal("Copy Fialed", err)
	// }
	// fmt.Println("Number of bytes writter", w)

	// 3 rules

	// 1 - A deferred function's arguments are evaluated when the defer statement is evaluated.
	// a()

	// 2 - Deferred function calls are executed in Last In First Out order after the surrounding function returns.
	// b()

	// 3 - Deferred functions may read and assign to the returning function's named return values.
	// fmt.Println(c())

	// panic
	// defer func() {
	r := recover()
	if r != nil {
		fmt.Println("Recovering from: ", r)
	}
	// }()
	fmt.Println(d())
	fmt.Println("Critical Info for the screen")
}
