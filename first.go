package main

import (
	"fmt"
	"reflect"
)

func main() {
	var price int = 100
	fmt.Println("The price is: ", price, "dollars")

	var taxRate float64 = 0.25
	fmt.Println("Tax rate is: ", taxRate)

	fmt.Println(reflect.TypeOf(price), reflect.TypeOf(taxRate))

	total := float64(price) * taxRate
	fmt.Println("Total: ", total, "TypeOf total", reflect.TypeOf(total))

	var availableFunds int = 120
	fmt.Println("availableFunds: ", availableFunds)

	newPrice := 100
	newTaxRate := 0.2
	var tax float64 = float64(newPrice) * newTaxRate

	fmt.Println("newPrice: ", newPrice, "tax", tax)
	
}
