package main

import (
	"fmt"
	"os"
)

type Entry struct {
	name     string
	lastname string
	tel      string
}

var data = []Entry{
	{"alex", "leon", "2109416471"},
	{"carlos", "lopez", "2109416871"},
	{"dian", "herrera", "2109416123"},
}

func search(key string) *Entry {
	for index, value := range data {
		if value.lastname == key {
			return &data[index]
		}

	}
	return nil
}

func list() {
	for _, value := range data {
		fmt.Println(value)
	}
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Debe especificar un comando valido")
		return
	}

	switch os.Args[1] {
	case "-s":
		if len(os.Args) != 3 {
			fmt.Println("uso: -s lastname")
			return
		}
		result := search(os.Args[2])
		if result == nil {
			fmt.Printf("El registro %v no fue encontrado", os.Args[2])
			return
		}
		fmt.Println(*result)
	case "-l":
		list()
	default:
		fmt.Println("Opción no valida")

	}
}

// go run main.go -s lastname, para buscar un registro
// go run main.go -l; para listar todos los registros
