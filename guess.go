package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)

	target := rand.Intn(100)

	//fmt.Println("Seconds:   ", seconds)
	//

	fmt.Println("I've chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")
	fmt.Println("Target:   ", target)

	for i := 10; i > 0; i-- {
		fmt.Print("Make a guess: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {

			log.Fatal(err)
		}
		number, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Enter the integer number")
			log.Fatal(err)
		}
		if number == target {
			fmt.Println("you guess right ")
			break
		} else {
			fmt.Println("you failed to guess")
			continue
		}

	}
	fmt.Println("x: ")
	x := 0
	x += 2
	fmt.Println(x)
	x -= 1
	fmt.Println(x)

	{
		for x := 1; x <= 3; x++ {
			fmt.Println(x)
		}
	}

	for x := 1; true; x++ {
		if x == 5 {
			break

		}
		fmt.Println(x)
	}

	for x := 1; x <= 3; x++ {
		y := x + 1
		fmt.Println("y inside cycle", y)
	}
	for x := 1; x <= 3; x++ {
		fmt.Println(x)
	}

	var z int
	for z = 0; z < 5; z++ {
		fmt.Println("z inside cycle", z)
	}
	fmt.Println("z outside cycle", z)

	for x := 1; x <= 3; x++ {
		fmt.Print(x)
	}
	for x := 3; x >= 1; x-- {
		fmt.Print(x)
	}
	for x := 2; x <= 3; x++ {
		fmt.Print(x)
	}

}
