package main

import (
	"contact-phone-book/options"
	"contact-phone-book/utils"
	"fmt"
	"os"
)

func main() {
	utils.GeneratesRecords()

	arguments := os.Args
	switch {
	case len(arguments) < 2 || len(arguments) > 3:
		fmt.Println("Debe especificar una opción, ejemplo -l|-s <arguments>")
		return
	case len(arguments) != 1:
		options.DetectInput(arguments[1])
	default:
		fmt.Println("Debe especificar una opción, ejemplo -l|-s <arguments>")
		return
	}

}
