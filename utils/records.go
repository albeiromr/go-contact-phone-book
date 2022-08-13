package utils

import (
	"contact-phone-book/models"
	"math/rand"
	"strconv"
	"time"
)

var Data []models.Person = []models.Person{}

//crea datos de prueba aleatorios
func GeneratesRecords() {
	for i := 0; i < 20; i++ {
		entry := models.Person{
			Name: generateName(10),
			Id:           generateStringId(),
		}
		Data = append(Data, entry)
	}
}

//genera un nombre random con un length dado
func generateName(length int) string {
	rand.Seed(time.Now().UnixNano())
	charset := "abcdefghijklmnopqrstuvwxyz"
	byteSlice := make([]byte, length)
	for i := 0; i < length; i++ {
		char := charset[rand.Intn(len(charset)-1)]
		byteSlice[i] = char
	}
	name := string(byteSlice)
	return name
}

//crea un string con números random entre 0 y 50
func generateStringId() string {
	return strconv.Itoa(rand.Intn(30))
}
