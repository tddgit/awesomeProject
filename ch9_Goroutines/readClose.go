package main

import "fmt"

func main() {
	willClose := make(chan int, 10)

	willClose <- -1
	willClose <- 0
	willClose <- 2

	<-willClose
	<-willClose
	<-willClose

	close(willClose)
	read1 := <-willClose
	fmt.Println(read1)
}

func f1(c chan int, x int) {
	fmt.Println(x)
}

func f2(c chan<- int, x int) {
	fmt.Println(x)
}

func f3(out chan<- int64, in <-chan int64) {
	fmt.Println(out)
	out <- in
}

func f4(out chan int64, in chan int64) {

}
