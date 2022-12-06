package main

import (
	"fmt"
	"sync"
)

func main() {
	char, num := make(chan bool), make(chan bool)

	wg := sync.WaitGroup{}
	go func() {

		for c := 'A'; c < 'Z'; c += 2 {
			char <- true
			fmt.Print(string(c))
			fmt.Print(string(c + 1))
			<-num

		}
		close(char)
	}()

	wg.Add(1)
	go func() {
		n := 1
		for {
			num <- true

			fmt.Print(n)
			fmt.Print(n + 1)
			n += 2

			_, ok := <-char
			if ok == false {
				break
			}
		}
		wg.Done()
	}()

	<-num

	wg.Wait()
	fmt.Print("\n")

}
