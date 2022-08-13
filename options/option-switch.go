package options

import (
	"fmt"
	"os"
)

//detecta si el usuario desea usar la opciós de list o de show, y redirige hacia las funciones de cada uno
func DetectInput(option string) {
	switch {
	case option == "-l" && len(os.Args) < 3:
		HandleListArgument()
	case option == "-l" && len(os.Args) > 2:
		fmt.Printf("Error: argumento (%v) invalido. uso -l \n", os.Args[2])
		return
	case option == "-s" && len(os.Args) < 3:
		fmt.Println("Debe especificar un argumento de busqueda valido. uso -s <numeric id>")
		return
	case option == "-s" && len(os.Args) > 2:
		HandleShowArgument(os.Args[2])
	default:
		fmt.Println("Debe especificar una opción valida, ejemplo -l|-s <arguments>")
		return
	}
}
