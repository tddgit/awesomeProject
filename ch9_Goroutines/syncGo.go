package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	n := flag.Int("n", 20, "Number of goroutines")
	//fmt.Println("n: ", reflect.TypeOf(n))
	flag.Parse()
	count := *n
	fmt.Printf("Going to create %d goroutines\n", count)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	fmt.Printf("%#v\n", waitGroup)

	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}
	waitGroup.Done()
	fmt.Println()
	fmt.Printf("%#v\n ", waitGroup)
	waitGroup.Wait()
	fmt.Println("\nExiting")

}
