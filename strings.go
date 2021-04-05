package main

import "fmt"

func main() {
	const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	s2 := "€£"
	fmt.Println(sLiteral, "    ", len(s2))
	fmt.Printf("(HEX) sLiteral: %x\n", sLiteral)
	fmt.Printf("(HEX) sLiteral: %c\n", sLiteral)
	fmt.Printf("(HEX) sLiteral: %s\n", sLiteral)
	fmt.Printf("sLiteral length: %d\n", len(sLiteral))

	for i := 0; i < len(sLiteral); i++ {
		fmt.Printf("%x ", sLiteral[i])
	}
	
}
