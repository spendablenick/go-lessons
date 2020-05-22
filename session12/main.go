package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func print10(n int) {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Number: ", n*10)
}

func urlSize(i int, in chan string, out chan string) {
	for url := range in {
		fmt.Println("Worker ", i, url)
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		out <- fmt.Sprintf("Worker %d, Url: %s, Result %d\n", i, url, len(body))
	}
}

func main() {
	urls := []string{
		"https://adjust.com",
		"https://example.com",
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://twitter.com",
		"https://blog.golang.org",
		"https://youtube.com",
		"https://vimeo.com",
		"https://codementor.io",
	}
	in := make(chan string)
	out := make(chan string)

	for i := 1; i <= 3; i++ {
		go urlSize(i, in, out)
	}

	var wg sync.WaitGroup

	go func() {
		for v := range out {
			fmt.Print("Result: ", v)
			wg.Done()
		}
	}()

	for _, v := range urls {
		in <- v
		wg.Add(1)
	}
	close(in)
	wg.Wait()
}
