package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year()

	fmt.Println("Current year", year)
	fmt.Println("Current time", now)
	fmt.Println("Current day", now.YearDay())

	broken := "G# r#cks!"
	fmt.Println(broken)
	replacer := strings.NewReplacer("#", "o")
	fmt.Println(reflect.TypeOf(replacer))
	fmt.Println(replacer.Replace(broken))
	
}
