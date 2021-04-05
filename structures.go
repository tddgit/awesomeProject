package main

import "fmt"

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	var s1 aStructure

	fmt.Println(s1)

	p1 := aStructure{"fmt", 12, -2}

	fmt.Println(p1)

	p1 = aStructure{person: "fmt", weight: -100, height: -200}

	fmt.Println(p1)

	type XYZ struct {
		X int
		Y int
		Z int
	}

	var s2 XYZ
	fmt.Println("s2.Y, s2.Z: ", s2.Y, s2.Z)
	
}
