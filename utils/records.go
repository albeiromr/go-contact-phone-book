package utils

import (
	"contact-phone-book/models"
	"math/rand"
	"strconv"
	"github.com/cip8/autoname"
)

var Data []models.Person = []models.Person{}

//crea datos de prueba aleatorios
func GeneratesRecords(){
	for i := 0; i < 20; i++ {
		entry := models.Person{
			CompleteName: autoname.Generate(),
			Id: generateStringId(),
		}
		Data = append(Data, entry)
	}
}

//crea un string con números random entre 0 y 50
func generateStringId() string {
	return strconv.Itoa(rand.Intn(30))
}
