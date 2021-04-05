// pass-fail inform about the passing exam
/*
Block commentary
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var packageVar = "package"

func main() {
	fmt.Println("=====================================================")
	var functionVar = "function"
	a := 10
	b, a := 20, 20
	a, c := 30, 30

	var a1, b1 float64
	var err error
	a1, err = strconv.ParseFloat("1.23", 64)
	b1, err = strconv.ParseFloat("2.56", 64)
	fmt.Println("a1, b1, err: ", a1, b1, err)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(b1), reflect.TypeOf(err))

	fmt.Println(a, b, c)
	fmt.Println(packageVar, functionVar)
	if true {
		conditionalVar := "conditional"
		fmt.Println(packageVar, functionVar, conditionalVar)
	}
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error is: ", err)
		//log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error is: ", err)
		fmt.Println("Please enter the valid float number")
		//log.Fatal(err)
	}

	bool, err1 := strconv.ParseBool("true")
	file, err2 := os.Open("first.go")

	fmt.Println(reflect.TypeOf(bool), reflect.TypeOf(file),
		reflect.TypeOf(err1), reflect.TypeOf(err2))

	if 1 < 2 {
		fmt.Println("1 < 2")
	}

	if true {
		fmt.Println("i'll be printed")

	}

	if false {
		fmt.Println("i  won't printed")
	}

	var status string
	if grade == 100 {
		status = "passing"
		fmt.Println("Perfect!", status)

	} else if grade >= 60 {
		status = "passing"

		fmt.Println("Quiet OK! You pass!", status)
	} else {
		status = "failed"
		fmt.Println("Very Bad. You fail!", status)
	}

	s := "\t formerly surruounded by space \n"
	fmt.Println(strings.TrimSpace(s))

	//response, err3 := http.Get("http://golang.org/")
	//fmt.Println(err3, response)

	if 1 == 1 {
		fmt.Println("i'll be printed 1 ==1 ")
	}
	if 1 >= 2 {
		fmt.Println("i'll not be printed")
	}

	if 1 > 2 {
		fmt.Println("i won't")
	}
	if 2 <= 2 {
		fmt.Println("i'll be printed 2 <=2")
	}

	if 1 < 2 {
		fmt.Println("i'll be printed 1 < 2")
	}

	if !true {
		fmt.Println("i won't be printed'")
	}

	if !false {
		fmt.Println("!false")

	}

	if true && true {
		fmt.Println("true && true")
	}
	if false && false {
		fmt.Println("i won't be printed")
	}
	if false && true {
		fmt.Println("i won't be printed")
	}

	if false || true {
		fmt.Println("false || true")
	}

	var languages = append([]string{}, "Espanol")
	fmt.Println(languages)

	{
		a := 100
		fmt.Println(a)
	}
	{
		a := 200
		fmt.Println(a)
	}

}
