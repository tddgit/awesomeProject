package main

import "fmt"

func main() {
	var width, height, area float64
	width = 4.2
	height = 4
	area = width * height
	fmt.Println(area/10.0, "liters needed")

	width = 5.2
	height = 4.67
	area = width * height
	fmt.Println(area/10.0, "liters needed")

	calculateArea(4.2, 4)
	calculateArea(5.2, 4.67)
	fmt.Println("Area", calculateArea(5, 5))

	fmt.Println("About one-third", 1.0/3.0)
	fmt.Printf("Printf: About 1/3: %0.4f, About 1/9: %0.6f", 1.0/3.0,
		1.0/9.0)
	fmt.Println()

	resultString := fmt.Sprintf("About one-third: %0.4f, About 1/9: %0.6f", 1.0/3.0, 1.0/9.0)
	fmt.Println("resultString", resultString)

}

func calculateArea(width, height float64) float64 {
	area := width * height
	fmt.Println(area/10.0, "liters needed")
	return area
}
