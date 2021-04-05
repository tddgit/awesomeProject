package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

var n = flag.Bool("n", false, "skipping the symbol of new row")
var sep = flag.String("s", " ", "delimiter")

func main() {

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)
	{
		var x, y int
		fmt.Println(&x == &x, &x == &y, &x == nil)
	}
	fmt.Println(p)
	fmt.Println(f() == f())

	z := 10
	incr(&z)
	fmt.Println(z)

	p1 := new(int)
	fmt.Println(reflect.ValueOf(p1), reflect.TypeOf(p1), p1, *p1)

	q1 := new(int)
	fmt.Println(p1 == q1)

	x = 1
	*p = 2000
	//person.name = bob
	//count[x] = count[x] * scale

	file, err := os.Open("first.go")
	i, j, k := 2, 3, 5
	i, j, k = 1, 2, 3
	fmt.Println(i, j, k, file, err)
	v := 1
	v--
	fmt.Println(v)

	medals := []string{"gold", "silver", "bronze"}
	medals[0] = "GOLD"
	fmt.Println(medals)
}

var p = f()

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
