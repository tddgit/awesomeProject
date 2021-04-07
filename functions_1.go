package main

import (
	"fmt"
	"math"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func f(i, j, k int, s, t string) {
	fmt.Println(i, j, k, s, t)
}

func f1(i int, j int, k int, s string, t string) {
	/* */
}

func add(x, y float64) float64      { return x + y }
func sub(x, y float64) (z float64)  { z = x - y; return }
func first(x int, y int) (int, int) { return x, y }
func zero(int, int) int             { return 0 }

func main() {
	fmt.Println(hypot(3, 4))
	fmt.Println(zero(5, 6))

	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", zero)
}
