package options

import (
	"contact-phone-book/logs"
	"contact-phone-book/models"
	"contact-phone-book/utils"
	"fmt"
)

func HandleShowArgument(argumento string) {

	var result []models.Person = []models.Person{}

	for _, value := range utils.Data {
		if value.Id == argumento {
			result = append(result, value)
		}
	}

	if len(result) <= 0 {
		logs.CreateLog(argumento)
		return
	}

	for _, value := range result {
		fmt.Printf("El Registro número (%v) fue encontrado y su valor es: %v\n", argumento, value)
	}

}
