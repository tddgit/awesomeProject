package main

import (
	"fmt"
	"time"
)

func main() {
	const pi = 3.14159
	const (
		e     = 2.718281
		newPi = 3.14159
	)
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	//fmt.Printf()

	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
	)

	fmt.Println(Sunday, Monday, Tuesday)

}
