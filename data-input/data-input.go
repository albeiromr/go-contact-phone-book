package datainput

import (
	"fmt"
	"regexp"
	"strconv"
)

type Person struct {
	Name     string
	lastName string
	id       string
}

func EnterPerson() {
	newPerson := Person{}

	var name string
	var lastName string
	var id string


	fmt.Println("Igrese el nombre de la persona")
	fmt.Scanln(&name)
	if !validateName(name) {
		panic("Error: formato de nombre invalido")
	}
	newPerson.Name = name

	fmt.Println("Igrese el apellido de la persona")
	fmt.Scanln(&lastName)
	if !validateName(lastName) {
		panic("Error: formato de apellido invalido")
	}
	newPerson.lastName = lastName

	fmt.Println("Igrese el id de la persona")
	fmt.Scanln(&id)
	if len(id) != 3 {
		panic("Error: formato de id invalido")
	}
	_, error := strconv.Atoi(id)
	if error != nil {
		panic("Error: formato de id invalido, solo se admiten números")
	}
	newPerson.id = id

	//TODO: se debe dar manejo a los datos obtenidos mandarlos al archivo csv
}

func validateName(name string) bool {
	data := []byte(name)
	regex := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return regex.Match(data)
}
