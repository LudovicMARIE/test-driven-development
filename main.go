package main

import (
	"fmt"

	"github.com/LudovicMARIE/tdd/roman"
)

func main() {
	fmt.Println("Hello World")
	fmt.Printf("10 --> %s\n", roman.ConvertIntToRomanNumeral(10))
	fmt.Printf("34 --> %s\n", roman.ConvertIntToRomanNumeral(34))
	fmt.Printf("1653 --> %s\n", roman.ConvertIntToRomanNumeral(1653))
}
