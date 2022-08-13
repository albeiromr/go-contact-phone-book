package main

import (
	"fmt"
	"contact-phone-book/utils"
)

func main(){
	utils.GeneratesRecords()
	for index, value := range utils.Data {
		fmt.Printf("el registro con el index %v es:  %v\n", index, value)
	}
}