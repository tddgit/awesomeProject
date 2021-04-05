package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x, y, x*y, x+y, x-y)

	x = 1 + 2i
	y = 1 + 2i

	fmt.Println(cmplx.Sqrt(-1))
	
}
