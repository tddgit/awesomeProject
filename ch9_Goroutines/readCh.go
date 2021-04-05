package main

import (
	"fmt"
	"time"
)

func writeToChan(c chan int, x int) {
	fmt.Println("1: ", x)
	c <- x
	close(c)
	fmt.Println("2: ", x)
}

func main() {
	c := make(chan int)
	go writeToChan(c, 10)
	time.Sleep(1 * time.Second)
	k := <-c
	fmt.Printf("Read: %d\n ", k)
	fmt.Printf("Read: %d\n ", <-c)
	time.Sleep(1 * time.Second)

	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed")
	}
}
