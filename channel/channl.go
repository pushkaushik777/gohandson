package main

import (
	"fmt"
	"time"
)

func receive(t chan string) (r chan string) {
	fmt.Println("Inside th Channel... ")
	t <- "done"
	return t
}

func main() {
	p := make(chan string, 1)
	go receive(p)
	n := <-p
	fmt.Println(n)

	a := make(chan string)
	b := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		a <- "From firstChannel"
	}()

	go func() {
		time.Sleep(6 * time.Second)
		b <- "From SecondChannel"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-a:
			fmt.Println("received", msg1)
		case msg2 := <-b:
			fmt.Println("received", msg2)
		case <-time.After(4 * time.Second):
			fmt.Println("TimLoggout ....")
		}
	}

}
