package main

import "fmt"

func function1(i int) int {
	return i + i
}

func function2(i int) int {
	return i * i
}

func funFun(f1 func(int) int, f2 func(int) int, v1 int, v2 int) (int, int) {
	return f1(v1), f2(v2)
}

func main() {
	fmt.Println(funFun(function1, function2, 123, 123))

	fmt.Println(funFun(
		func(i int) int { return i * i * i },
		func(i int) int { return i * i * i * i },
		123,
		123))

}
