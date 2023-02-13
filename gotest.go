package main

import (
	"fmt"
	"sync"
)

func PrintSomething(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(str)
}

func main() {

	var wg sync.WaitGroup //nil

	wg.Add(2)

	go PrintSomething("First", &wg)
	go PrintSomething("Third", &wg)

	wg.Wait()
}
