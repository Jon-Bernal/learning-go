package main

import (
	"fmt"

	"github.com/Jon-Bernal/learning-go/mascot"
	// pulled in a module from the internet here
	"rsc.io/quote"
)

func main() {
	fmt.Println(mascot.BestMascot())
	// pulled in a module from the internet here
	fmt.Println(quote.Go())
}