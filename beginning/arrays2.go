package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
)

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	a[2] = 5
	fmt.Println(a)

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf(" %d\n", v)
	}

	var q [3]int64 = [3]int64{1, 2, 3}
	var r [3]int64 = [3]int64{1, 2, 3}
	fmt.Println(q == r)

	w := [...]int{1, 2, 3, 4}
	fmt.Printf("%T\n", w)
	w = [4]int{1, 2, 3, 6}

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RUR
	)

	symbol := [...]string{USD: "$", EUR: "E", GBP: "P", RUR: "RUR"}
	fmt.Println(RUR, symbol[RUR])

	m := [...]int{9: -1}
	fmt.Println(m)
	{
		a := [2]int{1, 2}
		b := [...]int{1, 2}
		c := [2]int{1, 3}
		d := [3]int{1, 2}
		fmt.Println(a, b, c, d)
	}

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	arr := [32]byte{}
	fmt.Printf("%T\t%x\n", arr, arr)

	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April",
		5: "May", 6: "June", 7: "July", 8: "August", 9: "September",
		10: "October", 11: "November", 12: "December"}

	fmt.Printf("%T\t%v\n", months, months)

	Q2 := months[4:7]
	fmt.Printf("%T\t%v\n", Q2, Q2)
	fmt.Println(reflect.TypeOf(Q2), len(Q2), cap(Q2))
	Summer := months[6:9]
	fmt.Printf("%T\t%v\n", Summer, Summer)

	for _, month1 := range Q2 {
		for _, month2 := range Summer {
			if month1 == month2 {
				fmt.Printf("%s is in both slices\n", month1)
			}
		}
	}

}

func zero(ptr *[32]byte) {
	go
	for i := range ptr {
		ptr[i] = 0
	}
}

func zero1(prt *[32]byte) {
	*prt = [32]byte{}
}
