package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 123

	s := fmt.Sprintf("x=%x", x)
	fmt.Println(s)

	x, err := strconv.Atoi("123")
	y, err := strconv.ParseInt("123", 10, 0)

	fmt.Println(x, y, err)

	string := "hello, world"
	//c := string[len(string)]

	fmt.Println(len(string), string[0], string[11])

	fmt.Println(string[7:12])
	fmt.Println(string[:5], string[7:], string[:])
	fmt.Println("goodbye " + string[7:])

	s = "left leg"
	t := s
	s += ", right leg"

	fmt.Println(s, "t:", t)
	
}
