package main

import (
	"go-contact-phone-book/file-factory"
)

func main() {

	//creando base de datos en formato .csv
	csvFileError := filefactory.UseFileFactory("database", filefactory.CsvFile)
	if csvFileError != nil {
		panic(csvFileError)
	}

	//creando archivo de logs en formato .log
	logsFileError := filefactory.UseFileFactory("logs", filefactory.LogFile)
	if logsFileError != nil {
		panic(logsFileError)
	}

}
