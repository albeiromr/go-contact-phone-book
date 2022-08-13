package options

import (
	"contact-phone-book/utils"
	"fmt"
)

func HandleListArgument(){
	for index, value := range utils.Data {
		fmt.Printf("%v : %v \n", index, value)
	}
}