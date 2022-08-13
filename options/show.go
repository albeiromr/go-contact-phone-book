package options

import (
	"contact-phone-book/utils"
	"contact-phone-book/logs"
	"fmt"
)

func HandleShowArgument(argumento string){
	for index, value := range utils.Data {
		if value.Id == argumento {
			fmt.Printf("registro (%v) encontrado, el valor es: %v\n", argumento, utils.Data[index])
		} else {
			logs.CreateLog(argumento)
		}
	}
}


