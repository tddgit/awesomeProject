package main

import (
	"fmt"
	"time"
)

func main() {
	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[6] = "mi"
	fmt.Println(notes)
	fmt.Println(notes[0])
	fmt.Println(notes[1])

	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[4] = 5
	fmt.Println(numbers)

	var dates [3]time.Time
	dates[0] = time.Unix(124555344354, 0)
	dates[1] = time.Unix(124555344354, 0)
	dates[2] = time.Unix(124555344354, 0)
	fmt.Printf("%#v", dates)
	fmt.Println()
	var noteS [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Printf("%#v", noteS)
	fmt.Println()

	var primes [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%#v", primes)
	fmt.Println()

	var text [3]string = [3]string{
		"asdfasdfadfadf",
		"adfadfadsfadfadfa",
		"sfdasdfadsfadsfadf",
	}
	fmt.Printf("%#v", text)
	fmt.Println()
	{
		var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
		index := 1
		fmt.Println(index, ":", notes[index])
		index = 6
		fmt.Println(index, ":", notes[index])
		for i := 0; i < 7; i++ {
			fmt.Println(notes[i])
		}
		fmt.Println(len(notes))
		primes := [3]int{1, 2, 3}
		fmt.Println(len(primes))

	}

}
