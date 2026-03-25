package main

import (
	"fmt"

	"github.com/LudovicMARIE/tdd/roman"
	"github.com/LudovicMARIE/tdd/video_store"
)

func main() {
	fmt.Println("Hello World")
	fmt.Printf("10 --> %s\n", roman.ConvertIntToRomanNumeral(10))
	fmt.Printf("34 --> %s\n", roman.ConvertIntToRomanNumeral(34))
	fmt.Printf("1653 --> %s\n", roman.ConvertIntToRomanNumeral(1653))

	movieCollection := []video_store.Movie{
		{Title: "The Matrix", Year: 1999},
		{Title: "A beautiful mind", Year: 2001},
		{Title: "Intouchable", Year: 2011},
		{Title: "Forest Gump", Year: 1994},
	}

	moviesFound := video_store.FindByTitle("", movieCollection)

	fmt.Printf("Résultat de la recherche : %+v\n", moviesFound)
}
